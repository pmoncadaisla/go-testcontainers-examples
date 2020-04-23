# Go testcontainers examples

## Kafka

```
cd kafka/
go test
```

output:

```
2020/04/23 11:43:09 Starting container id: 97299fc9154e image: quay.io/testcontainers/ryuk:0.2.3
2020/04/23 11:43:09 Waiting for container id 97299fc9154e image: quay.io/testcontainers/ryuk:0.2.3
2020/04/23 11:43:09 Container is ready id: 97299fc9154e image: quay.io/testcontainers/ryuk:0.2.3
2020/04/23 11:43:09 Starting container id: 25c770de58f1 image: wurstmeister/zookeeper
2020/04/23 11:43:10 Waiting for container id 25c770de58f1 image: wurstmeister/zookeeper
2020/04/23 11:43:10 Container is ready id: 25c770de58f1 image: wurstmeister/zookeeper
2020/04/23 11:43:10 Starting container id: aa55d4d190d3 image: confluentinc/cp-kafka:5.2.1
2020/04/23 11:43:10 Container is ready id: aa55d4d190d3 image: confluentinc/cp-kafka:5.2.1
2020/04/23 11:43:17 Initializing new client
2020/04/23 11:43:17 client/metadata fetching metadata for all topics from broker :32888
2020/04/23 11:43:17 Connected to broker at :32888 (unregistered)
2020/04/23 11:43:17 client/brokers registered new broker #1 at localhost:32888
2020/04/23 11:43:17 Successfully initialized new client
2020/04/23 11:43:17 client/metadata fetching metadata for [topic] from broker :32888
2020/04/23 11:43:17 client/metadata found some partitions to be leaderless
2020/04/23 11:43:17 client/metadata retrying after 1000ms... (30 attempts remaining)
2020/04/23 11:43:18 client/metadata fetching metadata for all topics from broker :32888
2020/04/23 11:43:18 client/metadata fetching metadata for [topic] from broker :32888
2020/04/23 11:43:18 producer/broker/1 starting up
2020/04/23 11:43:18 producer/broker/1 state change to [open] on topic/0
2020/04/23 11:43:18 Connected to broker at localhost:32888 (registered as #1)
topic=topic	partition=0	offset=0
2020/04/23 11:43:18 Initializing new client
2020/04/23 11:43:18 client/metadata fetching metadata for all topics from broker :32888
2020/04/23 11:43:18 Connected to broker at :32888 (unregistered)
2020/04/23 11:43:18 client/brokers registered new broker #1 at localhost:32888
2020/04/23 11:43:18 Successfully initialized new client
2020/04/23 11:43:18 Connected to broker at localhost:32888 (registered as #1)
2020/04/23 11:43:18 consumer/broker/1 added subscription to topic/0
Partition:	0
Offset:	0
Key:
Value:	{ "some_json_data": 1 }

2020/04/23 11:43:18 consumer/broker/1 closed dead subscription to topic/0
2020/04/23 11:43:18 Closing Client
2020/04/23 11:43:18 Producer shutting down.
2020/04/23 11:43:18 Closing Client
2020/04/23 11:43:18 Closed connection to broker localhost:32888
2020/04/23 11:43:18 Closed connection to broker :32888
2020/04/23 11:43:18 producer/broker/1 input chan closed
2020/04/23 11:43:18 Closed connection to broker :32888
2020/04/23 11:43:18 producer/broker/1 shut down
2020/04/23 11:43:18 Closed connection to broker localhost:32888
PASS
ok  	github.com/pmoncadaisla/golang-testcontainers/kafka	9.966s
```