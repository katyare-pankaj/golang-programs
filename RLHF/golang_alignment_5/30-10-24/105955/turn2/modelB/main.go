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
	// Metric to store the count of code reviews processed at regular intervals.
	codeReviewsProcessed = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "code_reviews_processed_count",
			Help: "Number of code reviews processed at regular intervals.",
		},
		[]string{},
	)

	// Interval at which we want to sample and record data.
	samplingInterval = time.Second * 10
)

func main() {
	// Simulate code review processing and record metrics at a fixed interval.
	go func() {
		ticker := time.NewTicker(samplingInterval)
		defer ticker.Stop()

		for range ticker.C {
			processCodeReviews()
		}
	}()

	// Start the HTTP server to expose Prometheus metrics.
	fmt.Println("Prometheus metrics server listening on :8080")
	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func processCodeReviews() {
	// Simulate processing some code reviews.
	count := rand.Intn(10) + 1
	fmt.Printf("Processed %d code reviews during the interval.\n", count)

	// Record the sampled data point.
	codeReviewsProcessed.WithLabelValues().Set(float64(count))
}
