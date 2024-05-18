package config

import (
	"os"
)

type Config struct {
	ServerPort  string
	DatabaseURL string
	RabbitMQURL string
}

func LoadConfig() *Config {
	return &Config{
		ServerPort:  getEnv("SERVER_PORT", "8081"),
		DatabaseURL: getEnv("DATABASE_URL", "postgres://user:password@auction-db:5432/auctiondb?sslmode=disable"),
		RabbitMQURL: getEnv("RABBITMQ_URL", "amqp://guest:guest@rabbitmq:5672/"),
	}
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
