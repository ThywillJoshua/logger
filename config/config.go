package config

import (
	"os"
)

type Config struct {
	S3Bucket          string
	AWSRegion         string
	AWSAccessKeyID    string
	AWSSecretAccessKey string
}

func LoadConfig() *Config {
	return &Config{
		S3Bucket:          getEnv("S3_BUCKET_NAME", "default-bucket"),
		AWSRegion:         getEnv("AWS_REGION", "us-east-1"),
		AWSAccessKeyID:    getEnv("AWS_ACCESS_KEY_ID", ""),
		AWSSecretAccessKey: getEnv("AWS_SECRET_ACCESS_KEY", ""),
	}
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
