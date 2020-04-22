# Go testcontainers examples

## Kafka

```
cd kafka/
go test
```

output:

```
2020/04/22 22:16:30 Starting container id: 658f50ac3435 image: quay.io/testcontainers/ryuk:0.2.3
2020/04/22 22:16:30 Waiting for container id 658f50ac3435 image: quay.io/testcontainers/ryuk:0.2.3
2020/04/22 22:16:30 Container is ready id: 658f50ac3435 image: quay.io/testcontainers/ryuk:0.2.3
2020/04/22 22:16:30 Starting container id: 97fd3d7ffa9a image: wurstmeister/zookeeper
2020/04/22 22:16:30 Waiting for container id 97fd3d7ffa9a image: wurstmeister/zookeeper
2020/04/22 22:16:31 Container is ready id: 97fd3d7ffa9a image: wurstmeister/zookeeper
2020/04/22 22:16:31 Starting container id: f69c312d7f4a image: confluentinc/cp-kafka:5.2.1
2020/04/22 22:16:31 Container is ready id: f69c312d7f4a image: confluentinc/cp-kafka:5.2.1
2020/04/22 22:16:31 Initializing new client
2020/04/22 22:16:31 client/metadata fetching metadata for all topics from broker :32856
2020/04/22 22:16:31 Connected to broker at :32856 (unregistered)
2020/04/22 22:16:31 client/metadata got error from broker -1 while fetching metadata: EOF
2020/04/22 22:16:31 Closed connection to broker :32856
2020/04/22 22:16:31 client/metadata no available broker to send metadata request to
2020/04/22 22:16:31 client/brokers resurrecting 1 dead seed brokers
2020/04/22 22:16:31 client/metadata retrying after 1000ms... (30 attempts remaining)
2020/04/22 22:16:32 client/metadata fetching metadata for all topics from broker :32856
2020/04/22 22:16:32 Connected to broker at :32856 (unregistered)
2020/04/22 22:16:32 client/metadata got error from broker -1 while fetching metadata: EOF
2020/04/22 22:16:32 Closed connection to broker :32856
2020/04/22 22:16:32 client/metadata no available broker to send metadata request to
2020/04/22 22:16:32 client/brokers resurrecting 1 dead seed brokers
2020/04/22 22:16:32 client/metadata retrying after 1000ms... (29 attempts remaining)
2020/04/22 22:16:33 client/metadata fetching metadata for all topics from broker :32856
2020/04/22 22:16:33 Connected to broker at :32856 (unregistered)
2020/04/22 22:16:33 client/metadata got error from broker -1 while fetching metadata: EOF
2020/04/22 22:16:33 Closed connection to broker :32856
2020/04/22 22:16:33 client/metadata no available broker to send metadata request to
2020/04/22 22:16:33 client/brokers resurrecting 1 dead seed brokers
2020/04/22 22:16:33 client/metadata retrying after 1000ms... (28 attempts remaining)
2020/04/22 22:16:34 client/metadata fetching metadata for all topics from broker :32856
2020/04/22 22:16:34 Connected to broker at :32856 (unregistered)
2020/04/22 22:16:34 client/metadata got error from broker -1 while fetching metadata: EOF
2020/04/22 22:16:34 Closed connection to broker :32856
2020/04/22 22:16:34 client/metadata no available broker to send metadata request to
2020/04/22 22:16:34 client/brokers resurrecting 1 dead seed brokers
2020/04/22 22:16:34 client/metadata retrying after 1000ms... (27 attempts remaining)
2020/04/22 22:16:35 client/metadata fetching metadata for all topics from broker :32856
2020/04/22 22:16:35 Connected to broker at :32856 (unregistered)
2020/04/22 22:16:35 client/metadata got error from broker -1 while fetching metadata: EOF
2020/04/22 22:16:35 Closed connection to broker :32856
2020/04/22 22:16:35 client/metadata no available broker to send metadata request to
2020/04/22 22:16:35 client/brokers resurrecting 1 dead seed brokers
2020/04/22 22:16:35 client/metadata retrying after 1000ms... (26 attempts remaining)
2020/04/22 22:16:36 client/metadata fetching metadata for all topics from broker :32856
2020/04/22 22:16:36 Connected to broker at :32856 (unregistered)
2020/04/22 22:16:36 client/brokers registered new broker #1 at localhost:32856
2020/04/22 22:16:36 Successfully initialized new client
2020/04/22 22:16:36 client/metadata fetching metadata for [topic] from broker :32856
2020/04/22 22:16:36 client/metadata found some partitions to be leaderless
2020/04/22 22:16:36 client/metadata retrying after 1000ms... (30 attempts remaining)
2020/04/22 22:16:37 client/metadata fetching metadata for all topics from broker :32856
2020/04/22 22:16:37 client/metadata fetching metadata for [topic] from broker :32856
2020/04/22 22:16:37 producer/broker/1 starting up
2020/04/22 22:16:37 producer/broker/1 state change to [open] on topic/0
2020/04/22 22:16:37 Connected to broker at localhost:32856 (registered as #1)
topic=topic	partition=0	offset=0
2020/04/22 22:16:37 Initializing new client
2020/04/22 22:16:37 client/metadata fetching metadata for all topics from broker :32856
2020/04/22 22:16:37 Connected to broker at :32856 (unregistered)
2020/04/22 22:16:37 client/brokers registered new broker #1 at localhost:32856
2020/04/22 22:16:37 Successfully initialized new client
2020/04/22 22:16:37 Connected to broker at localhost:32856 (registered as #1)
2020/04/22 22:16:37 consumer/broker/1 added subscription to topic/0
Partition:	0
Offset:	0
Key:
Value:	{ "some_json_data": 1 }

2020/04/22 22:16:37 Closing Client
2020/04/22 22:16:37 Producer shutting down.
2020/04/22 22:16:37 Closing Client
2020/04/22 22:16:37 Closed connection to broker localhost:32856
2020/04/22 22:16:37 Closed connection to broker :32856
2020/04/22 22:16:37 producer/broker/1 input chan closed
2020/04/22 22:16:37 producer/broker/1 shut down
2020/04/22 22:16:37 Closed connection to broker :32856
2020/04/22 22:16:37 Closed connection to broker localhost:32856
2020/04/22 22:16:37 consumer/broker/1 disconnecting due to error processing FetchRequest: EOF
2020/04/22 22:16:37 kafka: error while consuming topic/0: EOF
PASS
ok  	github.com/pmoncadaisla/golang-testcontainers/kafka	7.944s
```