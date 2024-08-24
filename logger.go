package logger

import (
	"context"
	"encoding/json"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"

	loggerConfig "github.com/thywill/logger/config"
	"github.com/thywill/logger/logentry"
	"github.com/thywill/logger/s3"
)

type Logger struct {
	s3Client    s3.S3Client
	bucketName  string
}

func NewLogger(cfg *loggerConfig.Config) (*Logger, error) {
	awsConfig, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(cfg.AWSRegion),
		config.WithCredentialsProvider(
			credentials.NewStaticCredentialsProvider(cfg.AWSAccessKeyID, cfg.AWSSecretAccessKey, ""),
		),
	)
	if err != nil {
		return nil, err
	}
	client := s3.NewClient(awsConfig)
	return &Logger{
		s3Client:   client,
		bucketName: cfg.S3Bucket,
	}, nil
}

func (l *Logger) UploadLog(ctx context.Context, entry logentry.LogEntry) error {
	objectKey := entry.CorrelationID
	logData, err := json.Marshal(entry)
	if err != nil {
		return err
	}
	return s3.UploadLog(ctx, l.s3Client, l.bucketName, objectKey, logData)
}
