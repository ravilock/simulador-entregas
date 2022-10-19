package kafka

import (
	"fmt"
	"log"
	"os"

	confKafka "github.com/confluentinc/confluent-kafka-go/kafka"
)

type KafkaConsumer struct {
	MessageChannel chan *confKafka.Message
}

func NewKafkaConsumer(messageChannel chan *confKafka.Message) *KafkaConsumer {
  return &KafkaConsumer{messageChannel}
}

func (k *KafkaConsumer) Consume() {
  configMap := confKafka.ConfigMap{
    "bootstrap.servers": os.Getenv("KafkaBootstrapServers"),
    "group.id": os.Getenv("KafkaConsumerGroupId"),
  }

  consumer, err := confKafka.NewConsumer(&configMap)
  if err != nil {
    log.Fatalf("Error consuming kafka message: %q", err)
  }

  topics := []string{os.Getenv("KafkaReadTopic")}
  consumer.SubscribeTopics(topics, nil)
  fmt.Println("A Kafka Consumer Has Been Started")
  for {
    message, err := consumer.ReadMessage(-1)
    if err == nil {
      k.MessageChannel <- message
    }
  }
}
