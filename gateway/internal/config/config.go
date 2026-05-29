package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port            string
	UserServiceAddr string
	TaskServiceAddr string
}

func NewConfig() *Config {
	_ = godotenv.Load()
	return &Config{
		Port:            getEnv("GATEWAY_PORT", ":8080"),
		UserServiceAddr: getEnv("USER_SERVICE_ADDR", "localhost:50051"),
		TaskServiceAddr: getEnv("USER_SERVICE_ADDR", "localhost:50052"),
	}
}

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}
