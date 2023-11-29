package main

import (
	"encoding/json"
	"github.com/IBM/sarama"
)

// This is a simple example of a Kafka producer.
func main() {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	producer, err := sarama.NewSyncProducer([]string{"localhost:29092"}, config)
	if err != nil {
		panic(err)
	}

	jsonBody, _ := json.Marshal(map[string]string{"foo": "bar"})
	_, _, err = producer.SendMessage(&sarama.ProducerMessage{
		Topic:     "my-topic",
		Value:     sarama.ByteEncoder(jsonBody),
		Partition: int32(0),
	})
}
