FROM golang AS builder

# Download and unzip kafka
WORKDIR /tmp
RUN wget https://archive.apache.org/dist/kafka/{{.kafka_version}}/kafka_{{.scala_version}}-{{.kafka_version}}.tgz -q -O kfaka.tgz
RUN mkdir kafka
RUN tar -xzf kfaka.tgz -C kafka --strip-components 1

# Build the runner
COPY go.mod go.sum /app/
WORKDIR /app
RUN go mod download
COPY cmd/run /app/cmd/run
RUN go get ./cmd/run

ARG jdk_version
FROM openjdk:{{.jdk_version}}

COPY --from=builder /tmp/kafka /app
COPY --from=builder /go/bin/run /bin

ENV PATH /app/bin:$PATH
WORKDIR /app

# The `exec` here is used to prevent sh from traping os signal.
# Different docker base images have different versions of sh, they
# behaves different when handling child process.
CMD exec run

EXPOSE 9092
