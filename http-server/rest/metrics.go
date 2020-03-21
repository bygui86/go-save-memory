package rest

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	opsProcessed = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "total_processed_ops",
			Help: "Total number of processed operations",
		},
	)
)

func RegisterCustomMetrics() {
	prometheus.MustRegister(opsProcessed)
}

func IncreaseOpsProcessed() {
	go opsProcessed.Inc()
}
