package main

import (
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
}
