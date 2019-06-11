
## Example docker-compose file for local development

```yaml
services:
  kafka:
    image: ysmood/kafka:1.0.2
    environment:
      KAFKA_ADVERTISED_HOST_NAME: localhost # This required for mac-docker
```

## Example docker-compose file for gitlab-ci

```yaml
test:
  services:
    - name: ysmood/kafka:1.0.2
      alias: kafka
```