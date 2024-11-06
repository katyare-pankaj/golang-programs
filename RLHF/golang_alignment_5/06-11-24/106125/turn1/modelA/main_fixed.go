package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Feedback represents a single feedback entry
type Feedback struct {
	Rating  int
	Comment string
}

// VisualizationData holds the aggregated visualization data
type VisualizationData struct {
	RatingCount   map[int]int
	TotalComments int
	mu            sync.Mutex // Added mutex for safe concurrent access
}

// generateRandomFeedback generates a slice of random Feedback entries
func generateRandomFeedback(count int) []Feedback {
	feedback := make([]Feedback, 0, count)
	for i := 0; i < count; i++ {
		feedback = append(feedback, Feedback{
			Rating:  rand.Intn(5) + 1,
			Comment: "Random comment",
		})
	}
	return feedback
}

// processFeedbackOptimized processes a slice of Feedback entries and generates VisualizationData optimally
func processFeedbackOptimized(feedback []Feedback) *VisualizationData {
	var wg sync.WaitGroup
	data := &VisualizationData{
		RatingCount: make(map[int]int),
	}

	// Create a worker pool to process feedback concurrently
	const workerCount = 4
	workerPool := make(chan struct{}, workerCount)

	for _, entry := range feedback {
		wg.Add(1)
		workerPool <- struct{}{}

		go func(entry Feedback) {
			defer func() {
				wg.Done()
				<-workerPool
			}()

			// Process feedback entry
			data.mu.Lock() // Lock to safely update shared data
			data.RatingCount[entry.Rating]++
			data.TotalComments++
			data.mu.Unlock() // Unlock after processing
		}(entry)
	}

	wg.Wait()
	return data
}

func main() {
	// Generate random feedback data
	const feedbackCount = 100000
	feedback := generateRandomFeedback(feedbackCount)

	// Start the timer
	start := time.Now()

	// Process feedback optimally
	data := processFeedbackOptimized(feedback)

	// Calculate and print performance
	duration := time.Since(start)
	fmt.Printf("Processed %d feedback entries in %s\n", feedbackCount, duration)
	fmt.Printf("Visualization Data: %+v\n", data)
}
