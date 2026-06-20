package config

import (
	"fmt"
	"os"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	GRPCPort string

	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string

	JWTSecret     string
	JWTAccessTTL  time.Duration
	JWTRefreshTTL time.Duration
}

func NewConfig() *Config {
	_ = godotenv.Load("auth-service/.env")

	return &Config{
		GRPCPort: getEnv("GRPC_PORT", ":50051"),

		DBHost:     getEnv("DB_HOST", "localhost"),
		DBPort:     getEnv("DB_PORT", "5432"),
		DBUser:     getEnv("DB_USER", "postgres"),
		DBPassword: getEnv("DB_PASSWORD", "postgres"),
		DBName:     getEnv("DB_NAME", "postgres"),

		JWTSecret:     getEnv("JWT_SECRET", "supersecret"),
		JWTAccessTTL:  ttlDuration("JWT_ACCESS_TTL", "1m"),
		JWTRefreshTTL: ttlDuration("JWT_REFRESH_TTL", "30m"),
	}
}

func ttlDuration(ttl, defaultTtl string) time.Duration {
	dur, err := time.ParseDuration(getEnv(ttl, defaultTtl))

	if err != nil {
		dur, err = time.ParseDuration(defaultTtl)
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
