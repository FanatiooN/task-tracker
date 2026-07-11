package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	TaskServiceAddr string
	DailyReportCron string
}

func NewConfig() *Config {
	_ = godotenv.Load("scheduler/.env")

	return &Config{
		TaskServiceAddr: getEnv("TASK_SERVICE_ADDR", "task-service:50052"),
		DailyReportCron: getEnv("DAILY_REPORT_CRON", "0 9 * * *"),
	}
}

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}
