package main

import (
	"os"
	"syscall"

	kit "github.com/ysmood/gokit"
)

func main() {
	zookeeper := kit.Exec("zookeeper-server-start.sh", "config/zookeeper.properties")
	go zookeeper.Do()

	confPath := "config/server.properties"

	host := os.Getenv("KAFKA_ADVERTISED_HOST_NAME")
	if host != "" {
		conf, err := kit.ReadStringFile("config/server.properties")
		kit.E(err)

		confPath += ".env"
		kit.OutputFile(confPath, conf+"\nadvertised.host.name="+host, nil)
	}

	kafka := kit.Exec("kafka-server-start.sh", confPath)
	go kafka.Do()

	kit.WaitSignal(syscall.SIGTERM)

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
