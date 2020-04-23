package main

import (
	"context"
	"fmt"
	"os"
	"strings"
	"testing"
	"time"

	"log"

	"github.com/Shopify/sarama"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

const (
	kafkaImage                = "confluentinc/cp-kafka"
	confluentPlatformVersion  = "5.2.1"
	starterScriptContainerDir = "/tmp/"
	starterScriptName         = "testcontainers_start.sh"
	kafkaPort                 = "9093"
	ZookeeperPort             = "2181"
)

func TestNginxLatestReturn(t *testing.T) {
	ctx := context.Background()

	// Create docker network for comunicaiton between kafka and zookeeper
	networkName := "new-network"
	net, err := testcontainers.GenericNetwork(ctx, testcontainers.GenericNetworkRequest{
		NetworkRequest: testcontainers.NetworkRequest{
			Name:           networkName,
			CheckDuplicate: true,
		},
	})

	// If there is no error, it has been created successfuly, but if
	// there is an error the most common error is that already exists
	// so we continue
	if err == nil {
		defer net.Remove(ctx)
	}

	// Create zookeeper container and we wait until it is listening
	// at tcp/2181
	reqZookeeper := testcontainers.ContainerRequest{
		Name:         "zookeeper",
		Image:        "wurstmeister/zookeeper",
		ExposedPorts: []string{"2181/tcp"},
		Networks: []string{
			networkName,
		},
		WaitingFor: wait.ForListeningPort("2181"),
	}
	_, err = testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: reqZookeeper,
		Started:          true,
	})
	if err != nil {
		t.Error(err)
	}

	workingDir, err := os.Getwd()
	if err != nil {
		t.Error(err)
	}

	// Remove starter script from previous test executions
	starterScriptLocalPath := workingDir + "/" + starterScriptName
	os.Remove(starterScriptLocalPath)

	// Start script will be mounted at this path in the docker container
	starterScriptContainerPath := starterScriptContainerDir + starterScriptName

	req := testcontainers.ContainerRequest{
		Image: kafkaImage + ":" + confluentPlatformVersion,
		Name:  "kafka",

		// This is very tricky
		// We override container's startup command with a custom script
		// This script waits until a firstly missing starter script is found. Then executes it.
		// This script contains important configuration options wich must be filled after container has started (its IPaddress)
		Cmd:          []string{"sh", "-c", "while [ ! -f " + starterScriptContainerPath + " ]; do sleep 0.1; done; chmod +x " + starterScriptContainerPath + "; " + starterScriptContainerPath},
		ExposedPorts: []string{fmt.Sprintf("%s/tcp", kafkaPort)},
		Networks: []string{
			networkName,
		},
		Env: map[string]string{
			"KAFKA_LISTENERS":                        fmt.Sprintf("OUTSIDE://0.0.0.0:%s,INSIDE://0.0.0.0:9092", kafkaPort),
			"KAFKA_LISTENER_SECURITY_PROTOCOL_MAP":   "OUTSIDE:PLAINTEXT,INSIDE:PLAINTEXT",
			"KAFKA_INTER_BROKER_LISTENER_NAME":       "INSIDE",
			"KAFKA_BROKER_ID":                        "1",
			"KAFKA_OFFSETS_TOPIC_REPLICATION_FACTOR": "1",
			"KAFKA_OFFSETS_TOPIC_NUM_PARTITIONS":     "1",
			"KAFKA_LOG_FLUSH_INTERVAL_MESSAGES":      "65535",
			"KAFKA_GROUP_INITIAL_REBALANCE_DELAY_MS": "0",
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

	host, err := kafkaC.Host(ctx)
	if err != nil {
		t.Error(err)
	}

	ip, err := kafkaC.ContainerIP(ctx)
	if err != nil {
		t.Error(err)
	}

	port, err := kafkaC.MappedPort(ctx, kafkaPort)
	if err != nil {
		t.Error(err)
	}

	kafkaConnStr := fmt.Sprintf("%s:%s", ip, port.Port())

	// This is the tricky part, now that the container is running we can run kafka
	// knowing it's IP address
	commandStr := `#!/bin/bash 
	export KAFKA_ZOOKEEPER_CONNECT=zookeeper:2181
	export KAFKA_ADVERTISED_LISTENERS=INSIDE://%s:9092,OUTSIDE://%s:%s
	/etc/confluent/docker/run
	`
	command := fmt.Sprintf(commandStr, ip, host, port.Port())

	// Write startup script in the container
	kafkaC.Exec(ctx, []string{"sh", "-c", "echo '" + command + "' > " + starterScriptContainerPath})

	topic := "topic"
	value := "{ \"some_json_data\": 1 }"
	brokerList := kafkaConnStr
	logger := log.New(os.Stderr, "", log.LstdFlags)

	sarama.Logger = logger

	//// Producer

	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Return.Successes = true

	// Wait until kafka cluster is available
	config.Metadata.RefreshFrequency = 1 * time.Second
	config.Metadata.Retry.Backoff = 1 * time.Second
	config.Metadata.Retry.Max = 30

	config.Producer.Retry.Backoff = 1 * time.Second
	config.Producer.Retry.Max = 30

	config.Producer.Partitioner = sarama.NewHashPartitioner
	config.ClientID = "test-go"

	message := &sarama.ProducerMessage{Topic: topic, Partition: -1}
	message.Value = sarama.StringEncoder(value)

	producer, err := sarama.NewSyncProducer(strings.Split(brokerList, ","), config)
	if err != nil {
		printErrorAndExit(69, "Failed to open Kafka producer: %s", err)
	}
	defer func() {
		if err := producer.Close(); err != nil {
			logger.Println("Failed to close Kafka producer cleanly:", err)
		}
	}()

	partition, offset, err := producer.SendMessage(message)
	if err != nil {
		printErrorAndExit(69, "Failed to produce message: %s", err)
	} else {
		fmt.Printf("topic=%s\tpartition=%d\toffset=%d\n", topic, partition, offset)
	}

	c, err := sarama.NewConsumer(strings.Split(brokerList, ","), config)
	if err != nil {
		printErrorAndExit(69, "Failed to start consumer: %s", err)
	}
	initialOffset := sarama.OffsetOldest

	pc, err := c.ConsumePartition(topic, partition, initialOffset)
	if err != nil {
		printErrorAndExit(69, "Failed to start consumer for partition %d: %s", partition, err)
	}
	for msg := range pc.Messages() {
		fmt.Printf("Partition:\t%d\n", msg.Partition)
		fmt.Printf("Offset:\t%d\n", msg.Offset)
		fmt.Printf("Key:\t%s\n", string(msg.Key))
		fmt.Printf("Value:\t%s\n", string(msg.Value))
		fmt.Println()
		break
	}

	if err := c.Close(); err != nil {
		logger.Println("Failed to close consumer: ", err)
	}

}

func printErrorAndExit(code int, format string, values ...interface{}) {
	fmt.Fprintf(os.Stderr, "ERROR: %s\n", fmt.Sprintf(format, values...))
	fmt.Fprintln(os.Stderr)
	os.Exit(code)
}
