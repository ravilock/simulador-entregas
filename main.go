package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	appKafka "github.com/ravilock/simulador-entregas/app/kafka"
	infraKafka "github.com/ravilock/simulador-entregas/infrastructure/kafka"
)

func getEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error Loading Environment Variables")
	}
}

func main() {
  getEnv()
  messageChannel := make(chan *kafka.Message)
  consumer := infraKafka.NewKafkaConsumer(messageChannel)
  go consumer.Consume()
  for message := range messageChannel {
    fmt.Println(string(message.Value))
    go appKafka.Produce(message)
  }
}
