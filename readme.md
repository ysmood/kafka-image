
# Kafka Docker Image

Properly start and stop the process of kafka and zookeeper as a single node for deveopment use.

## Example docker-compose file for local development

```yaml
services:
  kafka:
    image: ysmood/kafka:1.0.3
    environment:
      KAFKA_ADVERTISED_HOST_NAME: localhost # This required for mac-docker
```

## Example docker-compose file for gitlab-ci

```yaml
test:
  services:
    - name: ysmood/kafka:1.0.3
      alias: kafka
```
