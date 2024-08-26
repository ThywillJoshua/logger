package logger

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"

	loggerConfig "github.com/thywilljoshua/logger/config"
	"github.com/thywilljoshua/logger/logentry"
	"github.com/thywilljoshua/logger/s3"
	"github.com/thywilljoshua/logger/utils"
)

type Logger struct {
	s3Client   s3.S3Client
	bucketName string
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
	objectKey := fmt.Sprintf("%s%s%s%s%s%s%s%s%s",
	    utils.RemoveSlashesAndConvertToString(entry.ServiceName)+"/",
		fmt.Sprintf("%02d/", entry.Timestamp.Year()),
		fmt.Sprintf("%02d/", entry.Timestamp.Month()),
		fmt.Sprintf("%02d/", entry.Timestamp.Day()),
		utils.RemoveSlashesAndConvertToString(entry.LogLevel)+"/",
		utils.RemoveSlashesAndConvertToString(entry.CorrelationID)+"-",
		utils.RemoveSlashesAndConvertToString(entry.HTTPMethod)+"-",
		utils.RemoveSlashesAndConvertToString(entry.URL)+"-",
		utils.RemoveSlashesAndConvertToString(entry.ResponseStatus),
	)

	logData, err := json.Marshal(entry)
	if err != nil {
		return err
	}
	return s3.UploadLog(ctx, l.s3Client, l.bucketName, objectKey, logData)
}