package kafka

import (
	"log"
	"os"

	ckafka "github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

type KafkaConsumer struct {
	messageChan chan *ckafka.Message
}

func NewKafkaConsumer(msgChan chan *ckafka.Message) *KafkaConsumer {
	return &KafkaConsumer{
		messageChan: msgChan,
	}
}

func (kc *KafkaConsumer) Consume(topics []string) {
	if !(len(topics) > 0) {
		return
	}

	c, err := ckafka.NewConsumer(&ckafka.ConfigMap{
		"bootstrap.servers":        os.Getenv("KAFKA_BOOTSTRAP_SERVERS"),
		"group.id":                 os.Getenv("KAFKA_CONSUMER_GROUP_ID"),
		"auto.offset.reset":        "latest",
		"enable.auto.offset.store": false,
		"enable.auto.commit":       false,
	})

	if err != nil {
		log.Fatalf("error to obtain kafka consumer. Details: '%s'", err)
	}

	c.SubscribeTopics(topics, nil)

	log.Printf("kafka consumer has been started with the provided list of topics: %v\n", topics)

	run := true

	for run {
		msg, err := c.ReadMessage(-1)
		if err == nil {
			stOffset, err := c.StoreOffsets([]ckafka.TopicPartition{
				{Topic: msg.TopicPartition.Topic, Partition: msg.TopicPartition.Partition, Offset: msg.TopicPartition.Offset + 1},
			})
			if err == nil {
				kc.messageChan <- msg
				log.Printf("offset store success: %v\n", stOffset)
			} else {
				log.Printf("offset store error from %v. Details: %v\n", msg.TopicPartition, err)
			}

		} else if !err.(ckafka.Error).IsTimeout() {
			log.Printf("consumer error: %v (%v)\n", err, msg)
		}
	}
}
