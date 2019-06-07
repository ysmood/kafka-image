FROM golang AS builder

# Build the runner
COPY . /run
WORKDIR /run
RUN go get -v .

# Download and unzip kafka
WORKDIR /tmp
ADD http://ftp.meisei-u.ac.jp/mirror/apache/dist/kafka/2.2.0/kafka_2.12-2.2.0.tgz kfaka.tgz
RUN tar -xvzf kfaka.tgz

FROM openjdk

COPY --from=builder /tmp/kafka_2.12-2.2.0 /app
COPY --from=builder /go/bin/run /bin

ENV PATH /app/bin:$PATH
WORKDIR /app
CMD run

EXPOSE 9092
