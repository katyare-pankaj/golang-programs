package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	requestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "request_duration_seconds",
			Help:    "Duration of requests",
			Buckets: prometheus.ExponentialBuckets(0.005, 2, 10),
		},
		[]string{"method", "endpoint"},
	)

	requestCount = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "request_count",
			Help: "Total number of requests",
		},
		[]string{"method", "endpoint", "status"},
	)
)

func init() {
	prometheus.MustRegister(requestDuration, requestCount)
}

// DummyRiskCalculation simulates a risk calculation function
func DummyRiskCalculation() error {
	time.Sleep(time.Duration(100 * time.Millisecond)) // Simulate some work
	return nil
}

func riskHandler(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	defer func() {
		method := r.Method
		endpoint := r.URL.Path
		status := fmt.Sprintf("%d", w.Header().Get("Status-Code"))

		requestDuration.WithLabelValues(method, endpoint).Observe(time.Since(start).Seconds())
		requestCount.WithLabelValues(method, endpoint, status).Inc()
	}()

	err := DummyRiskCalculation()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "Risk calculation completed")
}

func metricsHandler(w http.ResponseWriter, r *http.Request) {
	promhttp.Handler().ServeHTTP(w, r)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/risk", riskHandler)
	r.HandleFunc("/metrics", metricsHandler)

	fmt.Println("Server running on http://localhost:8080")
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		panic(err)
	}
}
