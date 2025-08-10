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

type JwtConfig struct {
	Secret   string
	ExpHours uint
}

type Config struct {
	Env  string
	Nats NatsConfig
	Http HttpConfig
	Jwt  JwtConfig
}

func Load() *Config {
	if err := godotenv.Load(); err != nil {
		fmt.Println("Error loading .env file. Loading default values")
	}

	natsPort, err := strconv.Atoi(getEnv("NATS_PORT", "4222"))
	if err != nil {
		panic("Cannot parse JWT_EXP_HOURS from env. Please set it to a valid integer value")
	}

	tokenExpHours, err := strconv.Atoi(getEnv("JWT_EXP_HOURS", "24"))
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
		Jwt: JwtConfig{
			Secret:   getEnv("JWT_SECRET", "my-super-secret"),
			ExpHours: uint(tokenExpHours),
		},
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
