package main

import (
	"expvar"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

var (
	codeReviewDuration = expvar.NewHistogram("code_review_duration_seconds")
	codeReviewsTotal   = expvar.NewInt("code_reviews_total")
	samplingInterval   = time.Second // Choose your desired sampling interval
	nextSampleTime     = time.Now()
)

func main() {
	// Simulate code review processing
	go func() {
		for {
			processCodeReview()
			time.Sleep(time.Millisecond * 500)
		}
	}()

	// Start a goroutine to sample and reset the expvar counters
	go func() {
		for {
			if time.Now().After(nextSampleTime) {
				sampleAndResetMetrics()
				nextSampleTime = nextSampleTime.Add(samplingInterval)
			}
			time.Sleep(time.Millisecond * 100)
		}
	}()

	// Start a HTTP server to expose expvar metrics
	fmt.Println("Expvar metrics server listening on :8080")
	expvar.HandleExpvar()
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func processCodeReview() {
	// Simulate code review processing time
	duration := time.Duration(rand.Intn(10)+1) * time.Second

	// Update the histogram without sampling
	codeReviewDuration.Add(float64(duration.Seconds()))
	codeReviewsTotal.Add(1)
}

func sampleAndResetMetrics() {
	// Perform sampling and reset counters here if needed
	// For example, you can calculate average duration and reset the histogram
	fmt.Println("Sampling metrics...")
	// Reset counters for the next sampling interval
	codeReviewDuration.Init()
	codeReviewsTotal.Set(0)
}
