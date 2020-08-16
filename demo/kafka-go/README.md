https://docs.confluent.io/current/quickstart/cos-docker-quickstart.html

# creat topic

```sh
docker-compose exec broker kafka-topics --create --bootstrap-server \
localhost:9092 --replication-factor 1 --partitions 1 --topic users
```

# go client
https://github.com/confluentinc/confluent-kafka-go
