
# Kafka Docker Image

Properly start and stop the process of kafka and zookeeper as a single node for development usage.
Other popular images like `wurstmeister/kafka` and `spotify/kafka` doesn't work well when being used for development.

## Install

[Releases](https://cloud.docker.com/u/ysmood/repository/docker/ysmood/kafka)

## Docker Example

```bash
docker run -p 9092:9092 -e KAFKA_ADVERTISED_HOST_NAME=localhost ysmood/kafka
```

## Example docker-compose file for local development

```yaml
services:
  kafka:
    image: ysmood/kafka
    environment:
      KAFKA_ADVERTISED_HOST_NAME: localhost # This required
```

## Example docker-compose file for gitlab-ci

```yaml
test:
  services:
    - name: ysmood/kafka
      alias: kafka
```
