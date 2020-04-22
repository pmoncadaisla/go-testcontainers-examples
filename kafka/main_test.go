package main

import (
	"context"
	"fmt"
	"testing"

	kafka "github.com/segmentio/kafka-go"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

func TestNginxLatestReturn(t *testing.T) {
	ctx := context.Background()

	networkName := "new-network"
	net, _ := testcontainers.GenericNetwork(ctx, testcontainers.GenericNetworkRequest{
		NetworkRequest: testcontainers.NetworkRequest{
			Name:           networkName,
			CheckDuplicate: true,
		},
	})
	defer net.Remove(ctx)

	reqZookeeper := testcontainers.ContainerRequest{
		Name:         "zookeeper",
		Image:        "wurstmeister/zookeeper",
		ExposedPorts: []string{"2181/tcp"},
		Networks: []string{
			networkName,
		},
		WaitingFor: wait.ForListeningPort("2181"),
	}
	zookeeperC, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: reqZookeeper,
		Started:          true,
	})
	if err != nil {
		t.Error(err)
	}

	_, err = zookeeperC.Host(ctx)
	if err != nil {
		t.Error(err)
	}

	req := testcontainers.ContainerRequest{
		Image:        "wurstmeister/kafka",
		Name:         "kafka",
		ExposedPorts: []string{"9092/tcp"},
		WaitingFor:   wait.ForListeningPort("9092"),
		Networks: []string{
			networkName,
		},
		Env: map[string]string{
			"KAFKA_ZOOKEEPER_CONNECT":    "zookeeper:2181",
			"KAFKA_CREATE_TOPICS":        "topic:1:1",
			"KAFKA_ADVERTISED_HOST_NAME": "localhost",
		},
	}

	kafkaC, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})

	if err != nil {
		t.Error(err)
	}
	defer kafkaC.Terminate(ctx)
	ip, err := kafkaC.Host(ctx)
	t.Logf("Kafka IP: %s", ip)
	if err != nil {
		t.Error(err)
	}
	port, err := kafkaC.MappedPort(ctx, "9092")
	if err != nil {
		t.Error(err)
	}
	t.Logf("Kafka Port: %s", port)

	kafkaConnStr := fmt.Sprintf("%s:%s", ip, port.Port())
	t.Logf("kafkaConnStr: %s", kafkaConnStr)

	// to produce messages
	topic := "topic"

	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{kafkaConnStr},
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	})

	w.WriteMessages(context.Background(),
		kafka.Message{
			Key:   []byte("Key-A"),
			Value: []byte("Hello World!"),
		},
		kafka.Message{
			Key:   []byte("Key-B"),
			Value: []byte("One!"),
		},
		kafka.Message{
			Key:   []byte("Key-C"),
			Value: []byte("Two!"),
		},
	)

	w.Close()

}
