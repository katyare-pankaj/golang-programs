package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	codeReviewDuration = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "code_review_duration_seconds",
			Help:    "Duration of code review processes in seconds.",
			Buckets: prometheus.ExponentialBuckets(0.01, 2, 10),
		},
		[]string{"status"},
	)

	codeReviewsTotal = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "code_reviews_total",
			Help: "Total number of code reviews.",
		},
		[]string{"status"},
	)
)

func main() {
	// Simulate code review processing
	go func() {
		for {
			processCodeReview()
			time.Sleep(time.Second * 5)
		}
	}()

	// Start HTTP server to expose Prometheus metrics
	fmt.Println("Prometheus metrics server listening on :8080")
	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func processCodeReview() {
	// Simulate code review processing time
	duration := time.Duration(rand.Intn(10)+1) * time.Second

	status := "success"
	if rand.Intn(2) == 0 {
		status = "failure"
	}

	// Record metrics
	codeReviewDuration.WithLabelValues(status).Observe(duration.Seconds())
	codeReviewsTotal.WithLabelValues(status).Inc()

	fmt.Printf("Code review processed in %s with status: %s\n", duration, status)
}
