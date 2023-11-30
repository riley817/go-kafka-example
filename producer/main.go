package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/riley817/go-kafka-example/producer/interface/router"
)

var writeLogRoutes router.WriteLogRoutes

func init() {
	writeLogRoutes = *router.NewWriteLogRoutes()

}

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	app.Post("/write/favorite", writeLogRoutes.RegisterFavorite)

	log.Fatal(app.Listen(":3000"))
}

// This is a simple example of a Kafka producer.
/*func main() {
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
}*/
