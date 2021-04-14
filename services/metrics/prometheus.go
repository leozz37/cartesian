package metrics

import (
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func IncCounter(counter prometheus.Counter) {
	counter.Inc()
}

func CreateCounter(name, help string) *prometheus.Counter {
	counter := promauto.NewCounter(prometheus.CounterOpts{
		Name: name,
		Help: help,
	})
	return &counter
}

func InitPrometheus() {
	log.Println("Starting Prometheus")
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":2112", nil)
}
