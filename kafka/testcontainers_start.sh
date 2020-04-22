#!/bin/bash 
	export KAFKA_ZOOKEEPER_CONNECT=zookeeper:2181
	export KAFKA_ADVERTISED_LISTENERS=INSIDE://:9092,OUTSIDE://localhost:32853
	/etc/confluent/docker/run
	