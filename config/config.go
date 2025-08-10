package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type NatsConfig struct {
	Host string
	Port int
}

type HttpConfig struct {
	Port string
}

type Config struct {
	Env  string
	Nats NatsConfig
	Http HttpConfig
}

func Load() *Config {
	if err := godotenv.Load(); err != nil {
		fmt.Println("Error loading .env file. Loading default values")
	}

	natsPort, err := strconv.Atoi(getEnv("NATS_PORT", "4222"))
	if err != nil {
		panic("Cannot parse JWT_EXP_HOURS from env. Please set it to a valid integer value")
	}

	return &Config{
		Env: getEnv("ENV", "dev"),
		Nats: NatsConfig{
			Host: getEnv("NATS_HOST", "localhost"),
			Port: natsPort,
		},
		Http: HttpConfig{
			Port: getEnv("HTTP_PORT", "8080"),
		},
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
