package consumers

import (
	"github.com/javiersrf/juris-workers/config"
	"github.com/streadway/amqp"
)

func Init(config config.Config) *amqp.Connection {
	connection, err := amqp.Dial(config.RabbitMQURL)

	if err != nil {
		panic(err)
	}

	return connection
}
