package cloudwatch

import "github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs"

type CloudWatchWriter struct {
	client        *cloudwatchlogs.Client
	logGroup      string
	logStream     string
	sequenceToken *string
}
