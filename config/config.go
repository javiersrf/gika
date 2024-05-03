package config

import "os"

func InitConfig() *Config {
	broker_url, exists := os.LookupEnv("RABBITMQ_URL")
	if !exists {
		broker_url = "amqp://guest:guest@localhost:5672/"
	}

	return &Config{
		RabbitMQURL: broker_url,
	}
}
