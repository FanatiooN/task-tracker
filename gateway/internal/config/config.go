package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	OAuth OAuthConfig

	Port            string
	UserServiceAddr string
	TaskServiceAddr string
	AuthServiceAddr string
	FrontendAddr    string
}

type OAuthConfig struct {
	GoogleClientID    string
	GoogleRedirectURI string
}

func NewConfig() *Config {
	_ = godotenv.Load("gateway/.env")

	return &Config{
		OAuth: OAuthConfig{
			GoogleClientID:    getEnv("OAUTH_GOOGLE_CLIENT_ID", "id"),
			GoogleRedirectURI: getEnv("OAUTH_GOOGLE_REDIRECT_URI", "http://localhost:3000/login/google"),
		},

		Port:            getEnv("GATEWAY_PORT", ":8080"),
		UserServiceAddr: getEnv("USER_SERVICE_ADDR", "localhost:50051"),
		TaskServiceAddr: getEnv("TASK_SERVICE_ADDR", "localhost:50052"),
		AuthServiceAddr: getEnv("AUTH_SERVICE_ADDR", "localhost:50053"),
		FrontendAddr:    getEnv("FRONTEND_ADDR", "http://localhost:3000"),
	}
}

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}
