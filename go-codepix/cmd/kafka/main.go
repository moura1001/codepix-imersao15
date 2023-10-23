package main

import (
	ckafka "github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/moura1001/codepix/infra/db"
	"github.com/moura1001/codepix/service/kafka"
)

func main() {
	database := db.GetDBConnection()
	producer := kafka.NewKafkaProducer()
	deliveryChan := make(chan ckafka.Event)

	go kafka.DeliveryReport(deliveryChan)
	kafka.NewKafkaProcessor(database, producer, deliveryChan).Consume()
}
