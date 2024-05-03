package main

import (
	"fmt"

	"github.com/javiersrf/juris-workers/config"
	"github.com/javiersrf/juris-workers/consumers"
	"github.com/javiersrf/juris-workers/handlers"
	"github.com/streadway/amqp"
)

func main() {
	fmt.Println("RabbitMQ in Golang: Getting started tutorial")
	config := config.InitConfig()
	conn := consumers.Init(*config)
	defer conn.Close()

	channel, err := conn.Channel()
	handlers.HandleError(err, "Error on opening Channel")

	defer channel.Close()
	fmt.Println("Successfully connected to RabbitMQ instance")

	forever := make(chan bool)
	worker1 := consumers.Consumer{
		Name:    "test",
		AutoAck: true,
		Callback: func(msg *amqp.Delivery) {
			fmt.Printf("Received Message : %s\n", msg.Body)
		},
	}
	worker2 := consumers.Consumer{
		Name:    "worker2",
		AutoAck: true,
		Callback: func(msg *amqp.Delivery) {
			fmt.Printf("Received Message : %s\n", msg.Body)
		},
	}

	go worker1.StartConsumer(channel)
	go worker2.StartConsumer(channel)
	<-forever

}
