package main

import (
	"fmt"
	"sync"
	"time"
)

// Function to simulate a processing stage in the pipeline
func processStage(in <-chan int, out chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for num := range in {
		time.Sleep(time.Duration(num) * time.Millisecond) // Simulate work
		out <- num * num
	}
}

func main() {
	const numStages = 3
	const numWorkers = 4
	const numItems = 10

	// Create channels for each stage
	var stages [numStages]chan int
	for i := 0; i < numStages; i++ {
		stages[i] = make(chan int, numItems*numWorkers)
	}

	// Create a wait group to manage Goroutines
	var wg sync.WaitGroup

	// Start workers for each stage
	for i := 0; i < numStages; i++ {
		for j := 0; j < numWorkers; j++ {
			wg.Add(1)
			go func(stage int) {
				defer wg.Done()
				if stage == 0 {
					// Input stage: generate numbers
					for i := 1; i <= numItems; i++ {
						stages[stage] <- i
					}
				} else {
					// Processing stage
					processStage(stages[stage-1], stages[stage], &wg)
				}
			}(i)
		}
	}

	// Output stage: collect results
	go func() {
		wg.Wait()
		close(stages[numStages-1])
	}()

	// Collect and print results
	results := make([]int, 0, numItems*numWorkers)
	for result := range stages[numStages-1] {
		results = append(results, result)
	}

	fmt.Println("Results:", results)
}
