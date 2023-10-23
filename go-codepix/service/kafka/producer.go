package kafka

import (
	"fmt"
	"log"
	"os"

	ckafka "github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

func NewKafkaProducer() *ckafka.Producer {
	p, err := ckafka.NewProducer(&ckafka.ConfigMap{
		"bootstrap.servers": os.Getenv("KAFKA_BOOTSTRAP_SERVERS"),
	})

	if err != nil {
		log.Fatalf("error to obtain kafka producer. Details: '%s'", err)
	}

	return p
}

func Publish(message, topic string, producer *ckafka.Producer, deliveryChan chan ckafka.Event) error {

	m := &ckafka.Message{
		TopicPartition: ckafka.TopicPartition{Topic: &topic, Partition: ckafka.PartitionAny},
		Value:          []byte(message),
	}

	err := producer.Produce(m, deliveryChan)
	if err != nil {
		return fmt.Errorf("error to enqueue the kafka message. Details: '%s'", err)
	}

	return nil
}

func DeliveryReport(deliveryChan chan ckafka.Event) {
	for e := range deliveryChan {
		switch ev := e.(type) {
		case *ckafka.Message:
			m := ev
			if m.TopicPartition.Error != nil {
				log.Printf("Delivery failed: %v\n", m.TopicPartition.Error)
			} else {
				log.Printf("Delivered message %s to: %s[%d]@%v\n", string(m.Value),
					*m.TopicPartition.Topic, m.TopicPartition.Partition, m.TopicPartition.Offset)
			}
		default:
			log.Printf("Ignored event: %s\n", ev)
		}
	}
}
