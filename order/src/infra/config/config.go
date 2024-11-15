package config

import (
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

func LoadConfig() *Config {
	cwd, _ := os.Getwd()
	envFilePath := filepath.Join(cwd, ".env")
	_ = godotenv.Load(envFilePath)

	cfg := &Config{
		ServerAddr:             os.Getenv("SERVER_ADDR"),
		GinMode:                os.Getenv("GIN_MODE"),
		LogGroupName:           os.Getenv("CLOUDWATCH_LOG_GROUP_NAME"),
		LogStreamName:          os.Getenv("LOG_STREAM_NAME"),
		PrometheusId:           os.Getenv("PROMETHEUS_WORKSPACE_ID"),
		GrafanaId:              os.Getenv("GRAFANA_WORKSPACE_ID"),
		CloudWatchLogGroupName: os.Getenv("CLOUDWATCH_LOG_GROUP_NAME"),
	}
	return cfg
}
