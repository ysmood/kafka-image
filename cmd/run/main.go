package main

import (
	"os"
	"syscall"

	kit "github.com/ysmood/gokit"
)

func main() {
	zookeeper := kit.Exec("zookeeper-server-start.sh", "config/zookeeper.properties")
	go zookeeper.Do()

	kafka := kit.Exec("kafka-server-start.sh", configKafkaViaEnvs("config/server.properties"))
	go kafka.Do()

	kit.WaitSignal(syscall.SIGTERM, os.Interrupt)

	// Must kill kafka first, or it will become a zombie process
	// If they are killed at the same time, a race condition may occur.
	// If zookeeper shutdown first, kafka will try to keep reconnecting it and ignore the terminal signal.
	// After a while docker will force kill everything, the pid file won't be removed by kafka process,
	// when restart kafka it will exit immediately because pid exists
	kit.KillTree(kafka.GetCmd().Process.Pid)
	// We have to wait kafka stops completely then kill zookeeper.
	kafka.GetCmd().Process.Wait()

	kit.KillTree(zookeeper.GetCmd().Process.Pid)
	zookeeper.GetCmd().Process.Wait()
}

func configKafkaViaEnvs(path string) string {
	host := os.Getenv("KAFKA_ADVERTISED_HOST_NAME")
	conf, err := kit.ReadStringFile("config/server.properties")
	kit.E(err)

	path += ".env"
	if host != "" {
		conf = conf + "\nadvertised.host.name=" + host
	}
	kit.OutputFile(path, conf, nil)
	return path
}
