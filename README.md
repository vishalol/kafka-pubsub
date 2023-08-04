# kafka-pubsub

### Starting the Kafka Cluster 

`MY_IP=your-ip docker-compose up -d`

#### Creating a topic

```bash
docker run --net=host --rm confluentinc/cp-kafka:3.0.1 kafka-topics --create --topic foo --partitions 4 --replication-factor 2 --if-not-exists --zookeeper 127.0.0.1:32181
```

#### Listen to a topic

```bash
kcat -C -b localhost:19092,localhost:29092,localhost:39092 -t foo -p 0
```

#### Publish to a topic 

```bash
echo 'publish to partition 0' | kcat -P -b 0.0.0.0:19092,0.0.0.0:29092,0.0.0.0:39092 -t foo -p 0
```


### Starting the Consumers

```bash 
cd worker && go run main.go
```


### Starting the Producer

```bash 
cd producer && go run main.go
```

### Publish message 

```bash
curl --location 'http://localhost:8080/api/message' \
--header 'Content-Type: application/json' \
--data '{
    "symbol": "VSL",
    "val": 101
}'
```
