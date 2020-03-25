package main

import (
	"github.com/ysmood/kit"
)

func main() {
	kit.Tasks().Add(
		kit.Task("build", "").Init(func(cmd kit.TaskCmd) func() {
			deploy := cmd.Flag("deploy", "").Short('d').Bool()
			test := cmd.Flag("test", "").Short('t').Bool()

			return func() {
				build(*test, *deploy)
			}
		}),
	).Do()
}

func build(test, deploy bool) {
	type target struct {
		kafka  string
		scala  string
		jdk    string
		latest bool
	}

	list := []target{
		target{
			kafka:  "2.2.1",
			scala:  "2.11",
			jdk:    "13",
			latest: true,
		},
		target{
			kafka: "0.11.0.3",
			scala: "2.11",
			jdk:   "8-jre",
		},
	}

	for _, t := range list {
		tag := "ysmood/kafka:" + t.kafka
		if t.latest {
			tag = "ysmood/kafka"
		}

		dockerfileTpl, err := kit.ReadString("Dockerfile")
		kit.E(err)
		dockerfile := kit.S(dockerfileTpl,
			"kafka_version", t.kafka,
			"scala_version", t.scala,
			"jdk_version", t.jdk,
		)
		path := "tmp/Dockerfile-" + t.kafka
		kit.OutputFile(path, dockerfile, nil)

		kit.Exec("docker", "build", "-f", path, "-t", tag, ".").MustDo()

		if test {
			kit.Exec(
				"docker", "run", "--rm", tag,
			).MustDo()

			return
		}

		if deploy {
			kit.Exec("docker", "push", tag).MustDo()
		}
	}

}
