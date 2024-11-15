package cloudwatch

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs/types"
)

// NewCloudWatchWriter creates a new CloudWatchWriter instance
func NewCloudWatchWriter(ctx context.Context, logGroup, logStream string, cfg aws.Config) (*CloudWatchWriter, error) {
	var err error

	client := cloudwatchlogs.NewFromConfig(cfg)

	// Create log stream if it doesn't exist
	_, err = client.CreateLogStream(ctx, &cloudwatchlogs.CreateLogStreamInput{
		LogGroupName:  &logGroup,
		LogStreamName: &logStream,
	})
	if err != nil && !isStreamAlreadyExistsError(err) {
		return nil, fmt.Errorf("failed to create log stream: %v", err)
	}

	return &CloudWatchWriter{
		client:    client,
		logGroup:  logGroup,
		logStream: logStream,
	}, nil
}

// Write sends log entries to AWS CloudWatch
func (cw *CloudWatchWriter) Write(p []byte) (n int, err error) {
	logEvent := string(p)
	logEventTimestamp := time.Now().UnixNano() / int64(time.Millisecond)

	// Create the log entry
	logEvents := []types.InputLogEvent{
		{
			Message:   &logEvent,
			Timestamp: aws.Int64(logEventTimestamp),
		},
	}

	// Push log to CloudWatch
	putLogEventsInput := &cloudwatchlogs.PutLogEventsInput{
		LogGroupName:  &cw.logGroup,
		LogStreamName: &cw.logStream,
		LogEvents:     logEvents,
		SequenceToken: cw.sequenceToken,
	}

	// Send the log event to CloudWatch
	output, err := cw.client.PutLogEvents(context.Background(), putLogEventsInput)
	if err != nil {
		return 0, err
	}

	// Update the sequence token for subsequent log entries
	cw.sequenceToken = output.NextSequenceToken

	return len(p), nil
}

func isStreamAlreadyExistsError(err error) bool {
	var awsErr *types.ResourceAlreadyExistsException
	return errors.As(err, &awsErr)
}
