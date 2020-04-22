# Go testcontainers examples

## Kafka

```
cd kafka/
go test
```

output:

```
2020/04/22 16:07:57 Starting container id: b056b4f28cc0 image: quay.io/testcontainers/ryuk:0.2.3
2020/04/22 16:07:57 Waiting for container id b056b4f28cc0 image: quay.io/testcontainers/ryuk:0.2.3
2020/04/22 16:07:57 Container is ready id: b056b4f28cc0 image: quay.io/testcontainers/ryuk:0.2.3
2020/04/22 16:07:57 Starting container id: 6fcb77ca29aa image: wurstmeister/zookeeper
2020/04/22 16:07:57 Waiting for container id 6fcb77ca29aa image: wurstmeister/zookeeper
2020/04/22 16:07:58 Container is ready id: 6fcb77ca29aa image: wurstmeister/zookeeper
2020/04/22 16:07:58 Starting container id: f6e9f1b1a9e3 image: wurstmeister/kafka
2020/04/22 16:08:03 Waiting for container id f6e9f1b1a9e3 image: wurstmeister/kafka
2020/04/22 16:08:05 Container is ready id: f6e9f1b1a9e3 image: wurstmeister/kafka
PASS
ok  	github.com/pmoncadaisla/golang-testcontainers/kafka	17.923s
```