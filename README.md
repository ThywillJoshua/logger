# Logger Library

This Go library allows you to log entries to an AWS S3 bucket. The library provides a simple interface to create and upload structured log entries.

## Prerequisites

- Go 1.18 or higher
- An AWS account with access to S3

## Installation

Install dependency:

```bash
git get https://github.com/thywilljoshua/logger.git
```

## Configuration

Ensure the following environment variables are set with your AWS credentials and S3 bucket details:

- `S3_BUCKET_NAME` - Name of the S3 bucket.
- `AWS_REGION` - AWS region for the S3 bucket.
- `AWS_ACCESS_KEY_ID` - Your AWS access key ID.
- `AWS_SECRET_ACCESS_KEY` - Your AWS secret access key.

You can set these in your terminal session:

```bash
export S3_BUCKET_NAME=my-bucket
export AWS_REGION=us-east-1
export AWS_ACCESS_KEY_ID=your-access-key-id
export AWS_SECRET_ACCESS_KEY=your-secret-access-key
```

## Usage

Here's how you can use the library in your Go application:

```go
package main

import (
	"context"
	"log"
	"time"

	"github.com/thywilljoshua/logger/config"
	"github.com/thywilljoshua/logger/logger"
	"github.com/thywilljoshua/logger/logentry"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Initialize the logger
	log, err := logger.NewLogger(cfg)
	if err != nil {
		log.Fatalf("Failed to create logger: %v", err)
	}

	// Create a log entry
	entry := logentry.LogEntry{
		Timestamp:      time.Now(),
		ServiceName:    "my-service",
		LogLevel:       "INFO",
		CorrelationID:  "12345",
		ErrorMessage:   "",
		Error:          "",
		UserContext:    "user123",
		HTTPMethod:     "GET",
		URL:            "/api/v1/resource",
		ResponseStatus: 200,
		Payload:        "{}",
		ExecutionTime:  123,
	}

	// Upload the log entry
	err = log.UploadLog(context.Background(), entry)
	if err != nil {
		log.Fatalf("Failed to upload log: %v", err)
	}

	log.Println("Log uploaded successfully")
}
```

## Summary

1. **Load Configuration**: `config.LoadConfig()` reads environment variables to set up the configuration.

2. **Initialize Logger**: `logger.NewLogger(cfg)` creates a logger instance with AWS configuration.

3. **Create Log Entry**: Construct a `logentry.LogEntry` with your log details.

4. **Upload Log Entry**: Use `log.UploadLog(ctx, entry)` to upload the log to S3.
