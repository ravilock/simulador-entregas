package kafka

import (
	"log"
	"os"

	confKafka "github.com/confluentinc/confluent-kafka-go/kafka"
)

func NewKafaProducer() *confKafka.Producer {
  configMap := &confKafka.ConfigMap{
    "bootstrap.servers": os.Getenv("KafkaBootstrapServers"),
    "group.id": os.Getenv("KafkaConsumerGroupId"),
  }
  producer, err := confKafka.NewProducer(configMap)
  if err != nil {
    log.Fatalf("Error producing kafka message: %q", err)
  }
  return producer
}

func Publish(message, topic string, producer *confKafka.Producer) error {
  kafkaMessage := &confKafka.Message{
    TopicPartition: confKafka.TopicPartition{Topic: &topic, Partition: confKafka.PartitionAny},
    Value: []byte(message),
  }

  err := producer.Produce(kafkaMessage, nil)
  if err != nil {
    return err
  }

  return nil
}
