package config

import (
	"fmt"
	"os"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	OAuth OAuthConfig

	GRPCPort string

	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string

	UserServiceAddr string
	KafkaBrokerAddr string

	JWTSecret     string
	JWTAccessTTL  time.Duration
	JWTRefreshTTL time.Duration
}

type OAuthConfig struct {
	GoogleClientID     string
	GoogleClientSecret string
	GoogleRedirectURI  string

	TelegramClientID     string
	TelegramClientSecret string
}

func NewConfig() *Config {
	_ = godotenv.Load("auth-service/.env")

	return &Config{
		OAuth: OAuthConfig{
			GoogleClientID:     getEnv("OAUTH_GOOGLE_CLIENT_ID", "id"),
			GoogleClientSecret: getEnv("OAUTH_GOOGLE_CLIENT_SECRET", "secret"),
			GoogleRedirectURI:  getEnv("OAUTH_GOOGLE_REDIRECT_URI", "http://localhost:3000/login/google"),

			TelegramClientID:     getEnv("OAUTH_TELEGRAM_CLIENT_ID", "id"),
			TelegramClientSecret: getEnv("OAUTH_TELEGRAM_CLIENT_SECRET", "secret"),
		},

		GRPCPort: getEnv("GRPC_PORT", ":50053"),

		DBHost:     getEnv("DB_HOST", "localhost"),
		DBPort:     getEnv("DB_PORT", "5432"),
		DBUser:     getEnv("DB_USER", "postgres"),
		DBPassword: getEnv("DB_PASSWORD", "postgres"),
		DBName:     getEnv("DB_NAME", "postgres"),

		UserServiceAddr: getEnv("USER_SERVICE_ADDR", "localhost:50051"),
		KafkaBrokerAddr: getEnv("KAFKA_BROKER_ADDR", "localhost:9092"),

		JWTSecret:     getEnv("JWT_SECRET", "supersecret"),
		JWTAccessTTL:  ttlDuration("JWT_ACCESS_TTL", "15m"),
		JWTRefreshTTL: ttlDuration("JWT_REFRESH_TTL", "30m"),
	}
}

func ttlDuration(ttl, defaultTtl string) time.Duration {
	dur, err := time.ParseDuration(getEnv(ttl, defaultTtl))
	if err != nil {
		dur, _ = time.ParseDuration(defaultTtl)
	}

	return dur
}

func (c *Config) DatabaseURL() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		c.DBUser, c.DBPassword, c.DBHost, c.DBPort, c.DBName)
}

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}
