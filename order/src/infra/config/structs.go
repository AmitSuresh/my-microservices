package config

type Config struct {
	ServerAddr             string
	GinMode                string
	LogGroupName           string
	LogStreamName          string
	PrometheusId           string
	GrafanaId              string
	CloudWatchLogGroupName string
}
