package main

import (
	"os"
	"syscall"

	kit "github.com/ysmood/gokit"
)

func main() {
	zookeeper := kit.Exec("zookeeper-server-start.sh", "config/zookeeper.properties")
	go zookeeper.Do()

	host := os.Getenv("KAFKA_ADVERTISED_HOST_NAME")
	if host == "" {
		host = "localhost"
	}
	conf, err := kit.ReadStringFile("config/server.properties")
	kit.E(err)

	confPath := "config/server.properties.env"
	kit.OutputFile(confPath, conf+"\nadvertised.host.name="+host, nil)

	kafka := kit.Exec("kafka-server-start.sh", confPath)
	go kafka.Do()

	kit.WaitSignal(syscall.SIGTERM)

	kit.KillTree(kafka.GetCmd().Process.Pid)
	kafka.GetCmd().Process.Wait()

	kit.KillTree(zookeeper.GetCmd().Process.Pid)
	zookeeper.GetCmd().Process.Wait()
}
