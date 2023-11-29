package main

import (
	"encoding/json"
	"fmt"
	"github.com/IBM/sarama"
)

func main() {
	config := sarama.NewConfig()

	client, err := sarama.NewClient([]string{"localhost:29092"}, config)
	if err != nil {
		panic(err)
	}

	consumer, err := sarama.NewConsumerFromClient(client)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := consumer.Close(); err != nil {
			panic(err)
		}
	}()

	partitions, err := consumer.ConsumePartition("my-topic", int32(0), 0)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := partitions.Close(); err != nil {
			panic(err)
		}
	}()

	for {
		select {
		case msg := <-partitions.Messages():
			fmt.Println("Consumed -> ")
			fmt.Println(string(msg.Key))
			fmt.Println(string(msg.Value))
		}
	}

	for msg := range partitions.Messages() {
		var data map[string]string

		if err := json.Unmarshal(msg.Value, &data); err != nil {
			panic(err)
		}

		for key, val := range data {
			fmt.Print("Key: ", key, " Value: ", val, "\n")
		}
	}
}
