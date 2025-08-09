package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type NatsConfig struct {
	Host string
	Port string
}

type Config struct {
	Env  string
	Nats NatsConfig
}

func Load() *Config {
	if err := godotenv.Load(); err != nil {
		fmt.Println("Error loading .env file. Loading default values")
	}

	return &Config{
		Env: getEnv("ENV", "dev"),
		Nats: NatsConfig{
			Host: getEnv("NATS_HOST", "localhost"),
			Port: getEnv("NATS_PORT", "4222"),
		},
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
