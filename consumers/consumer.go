package consumers

import (
	"log"

	"github.com/javiersrf/gika/handlers"
	"github.com/streadway/amqp"
)

type CallbackConsumer func(msg *amqp.Delivery)

type Consumer struct {
	Name     string
	AutoAck  bool
	Callback CallbackConsumer
}

func (consumer Consumer) StartConsumer(channel *amqp.Channel) {
	queue, err := channel.QueueDeclare(consumer.Name, true, false, false, false, nil)

	handlers.HandleError(err, "Could not declare `add` queue")

	msgs, err := channel.Consume(
		queue.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		panic(err)
	}
	go func() {
		for msg := range msgs {
			log.Printf("\n\nReceived Message on consumer %s:\n", consumer.Name)
			consumer.Callback(&msg)

		}
	}()
	log.Printf("\n\n[x] Waiting for new messages on consumer %s\nTo exit press CTRL+C", consumer.Name)
}
