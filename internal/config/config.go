package config

import (
	"log/slog"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	// Server configuration
	HTTPPort string

	// Database configuration
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassword string
	DbName     string
	DbSSLMode  string

	// External services
	GeoServiceGRPCHost string

	// Kafka configuration
	KafkaHost                 string
	KafkaConsumerGroup        string
	KafkaBasketConfirmedTopic string
	KafkaOrderChangedTopic    string
}

func New() Config {
	return Config{
		HTTPPort:                  getEnv("HTTP_PORT"),
		DbHost:                    getEnv("DB_HOST"),
		DbPort:                    getEnv("DB_PORT"),
		DbUser:                    getEnv("DB_USER"),
		DbPassword:                getEnv("DB_PASSWORD"),
		DbName:                    getEnv("DB_NAME"),
		DbSSLMode:                 getEnv("DB_SSLMODE"),
		GeoServiceGRPCHost:        getEnv("GEO_SERVICE_GRPC_HOST"),
		KafkaHost:                 getEnv("KAFKA_HOST"),
		KafkaConsumerGroup:        getEnv("KAFKA_CONSUMER_GROUP"),
		KafkaBasketConfirmedTopic: getEnv("KAFKA_BASKET_CONFIRMED_TOPIC"),
		KafkaOrderChangedTopic:    getEnv("KAFKA_ORDER_CHANGED_TOPIC"),
	}
}

func getEnv(key string) string {
	if err := godotenv.Load(".env"); err != nil {
		slog.Error("Error loading .env file", "error", err)
		os.Exit(1)
	}

	return os.Getenv(key)
}
