FROM golang AS builder

# Build the runner
COPY . /run
WORKDIR /run
RUN go get -v .

# Download and unzip kafka
ARG KAFKA_VERSION=2.2.1
ARG SCALA_VERSION=2.11
WORKDIR /tmp
ADD https://archive.apache.org/dist/kafka/${KAFKA_VERSION}/kafka_${SCALA_VERSION}-${KAFKA_VERSION}.tgz kfaka.tgz
RUN mkdir kafka
RUN tar -xvzf kfaka.tgz -C kafka --strip-components 1

FROM openjdk:12

COPY --from=builder /tmp/kafka /app
COPY --from=builder /go/bin/run /bin

ENV PATH /app/bin:$PATH
WORKDIR /app
CMD run

EXPOSE 9092
