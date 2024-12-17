package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Data represents the type of data being processed
type Data struct {
	ID int
}

// Stage represents a processing stage in the pipeline
type Stage func(chan Data, chan Data, context.Context)

// Pipeline processes data through multiple stages
func Pipeline(stages []Stage, ctx context.Context) {
	var wg sync.WaitGroup

	channels := make([]chan Data, len(stages)+1)
	for i := range channels {
		channels[i] = make(chan Data, 100) // Buffered channel
	}

	// Start each stage as a goroutine
	for i, stage := range stages {
		wg.Add(1)
		go func(stage Stage, in chan Data, out chan Data, i int, ctx context.Context) {
			defer wg.Done()
			for {
				select {
				case data, ok := <-in:
					if !ok {
						close(out)
						return
					}
					stage(data, out, ctx)
				case <-ctx.Done():
					close(out)
					return
				}
			}
		}(stage, channels[i], channels[i+1], i, ctx)
	}

	// Simulate data production
	go func() {
		for i := 0; i < 10000; i++ {
			data := Data{ID: i}
			select {
			case channels[0] <- data:
			case <-ctx.Done():
				return
			}
		}
		close(channels[0])
	}()

	// Wait for all stages to complete
	wg.Wait()
	fmt.Println("Pipeline processing complete.")
}

// Example processing stages
func Stage1(in chan Data, out chan Data, ctx context.Context) {
	for {
		select {
		case data, ok := <-in:
			if !ok {
				return
			}
			// Simulate work
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(10)))
			fmt.Printf("Stage 1 processed Data ID: %d\n", data.ID)
			out <- data
		case <-ctx.Done():
			return
		}
	}
}

func Stage2(in chan Data, out chan Data, ctx context.Context) {
	for {
		select {
		case data, ok := <-in:
			if !ok {
				return
			}
			// Simulate work
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(10)))
			fmt.Printf("Stage 2 processed Data ID: %d\n", data.ID)
			out <- data
		case <-ctx.Done():
			return
		}
	}
}

func main() {
	// Set up context for cancellation
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	stages := []Stage{Stage1, Stage2}

	// Start the pipeline
	Pipeline(stages, ctx)
}
