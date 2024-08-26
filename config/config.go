package config

type Config struct {
	S3Bucket          string
	AWSRegion         string
	AWSAccessKeyID    string
	AWSSecretAccessKey string
}

func LoadConfig(s3Bucket, awsRegion, awsAccessKeyID, awsSecretAccessKey string) *Config {
	return &Config{
		S3Bucket:          s3Bucket,
		AWSRegion:         awsRegion,
		AWSAccessKeyID:    awsAccessKeyID,
		AWSSecretAccessKey: awsSecretAccessKey,
	}
}


