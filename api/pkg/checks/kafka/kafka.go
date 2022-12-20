package kafka

import (
	"fmt"

	"github.com/Shopify/sarama"
)

func KafkaHealthCheck(kafkaConn string) string {

	producer, err := sarama.NewSyncProducer([]string{kafkaConn + ":9092"}, nil)
	if err != nil {
		return ("Unhealthy")
	}
	defer producer.Close()

	// Create a new message and send it to the "my-topic" topic
	message := &sarama.ProducerMessage{
		Topic: "health",
		Value: sarama.StringEncoder("Healthy"),
	}
	partition, offset, err := producer.SendMessage(message)
	if err != nil {
		return ("Unhealthy")
	} else {
		fmt.Println(partition, offset)
		return ("Healthy")
	}
}
