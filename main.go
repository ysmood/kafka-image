package main

import (
	"os"

	kit "github.com/ysmood/gokit"
)

func main() {
	go kit.Exec("zookeeper-server-start.sh", "config/zookeeper.properties").MustDo()

	host := os.Getenv("KAFKA_ADVERTISED_HOST_NAME")
	if host == "" {
		host = "localhost"
	}
	conf, err := kit.ReadStringFile("config/server.properties")
	kit.E(err)

	confPath := "config/server.properties.env"
	kit.OutputFile(confPath, conf+"\nadvertised.host.name="+host, nil)

	kit.Exec("kafka-server-start.sh", confPath).MustDo()
}
