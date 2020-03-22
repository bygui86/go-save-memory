package rest

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	getRequests = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "get_requests_total",
			Help: "Total number of GET requests",
		},
	)
	postRequests = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "post_requests_total",
			Help: "Total number of POST requests",
		},
	)
	putRequests = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "put_requests_total",
			Help: "Total number of PUT requests",
		},
	)
	deleteRequests = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "delete_requests_total",
			Help: "Total number of DELETE requests",
		},
	)
)

func RegisterCustomMetrics() {
	prometheus.MustRegister(getRequests)
	prometheus.MustRegister(postRequests)
	prometheus.MustRegister(putRequests)
	prometheus.MustRegister(deleteRequests)
}

func IncreaseGetRequests() {
	go getRequests.Inc()
}

func IncreasePostRequests() {
	go postRequests.Inc()
}

func IncreasePutRequests() {
	go putRequests.Inc()
}

func IncreaseDeleteRequests() {
	go deleteRequests.Inc()
}
