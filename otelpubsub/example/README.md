# PubSub Instrumentation Example

Demonstration of a subscriber receiving messages from a GCP topic

Run the Dockerized subscriber app using the command below: 

```bash
docker-compose up subscriber
```

Publish a message by starting the publisher app:

```bash
docker-compose up publisher
```

Shutdown the example services by running:

```bash
docker-compose down
```
