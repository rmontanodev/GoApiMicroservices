package config

import (
	"os"
)

type Config struct {
	ServerPort         string
	DatabaseURL        string
	RabbitMQURL        string
	QUEUE_USER_CREATED string
}

func LoadConfig() *Config {
	return &Config{
		ServerPort:         getEnv("SERVER_PORT", "8081"),
		DatabaseURL:        getEnv("DATABASE_URL", "postgres://user:password@auction-db:5432/auctiondb?sslmode=disable"),
		RabbitMQURL:        getEnv("RABBITMQ_URL", "amqp://guest:guest@rabbitmq:5672/"),
		QUEUE_USER_CREATED: getEnv("QUEUE_USER_CREATED", "user_created"),
	}
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
