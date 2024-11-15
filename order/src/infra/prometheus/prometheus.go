package prometheus

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	HttpRequests = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "gin_http_requests_total",
			Help: "Total number of HTTP requests processed",
		},
		[]string{"method", "route"},
	)

	HttpDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "gin_http_request_duration_seconds",
			Help:    "Histogram of HTTP request durations in seconds",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"method", "route"},
	)
	Uptime = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "server_uptime_seconds",
		Help: "Number of seconds the server has been running",
	})
)

func InitializePrometheus() {
	prometheus.MustRegister(HttpRequests)
	prometheus.MustRegister(HttpDuration)
	prometheus.MustRegister(Uptime)
}
