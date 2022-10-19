package kafka

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	route2 "github.com/ravilock/simulador-entregas/app/route"
	infraKafka "github.com/ravilock/simulador-entregas/infrastructure/kafka"
)

func Produce(receivedMessage *kafka.Message) {
	producer := infraKafka.NewKafaProducer()
	route := route2.NewRoute()
	json.Unmarshal(receivedMessage.Value, route)
	route.LoadPositions()
	positions, err := route.ExportJSONPositions()
	if err != nil {
		log.Println(err.Error())
	}

	for _, position := range positions {
    fmt.Println(position)
		infraKafka.Publish(position, os.Getenv("KafkaProduceTopic"), producer)
		time.Sleep(time.Millisecond * 500)
	}
}
