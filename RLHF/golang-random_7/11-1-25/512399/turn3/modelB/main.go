package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type DataPoint struct {
	Timestamp time.Time
	Value     float64
}

func generateData(ctx context.Context, dataCh chan<- DataPoint, id int, rate time.Duration) {
	defer close(dataCh)
	for {
		select {
		case <-ctx.Done():
			return
		default:
			dataPoint := DataPoint{
				Timestamp: time.Now(),
				Value:     rand.Float64() * 100,
			}
			dataCh <- dataPoint
			time.Sleep(rate)
		}
	}
}

func processData(dataCh <-chan DataPoint, processedCh chan<- DataPoint) {
	for dataPoint := range dataCh {
		// Simulate some processing work
		time.Sleep(time.Millisecond * 50)
		processedCh <- dataPoint
	}
	close(processedCh)
}

func visualizeData(processedCh <-chan DataPoint) {
	for dataPoint := range processedCh {
		// Simulate visualization update
		fmt.Printf("Visualizing data: %+v\n", dataPoint)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	numStreams := 3
	dataChannels := make([]chan DataPoint, numStreams)
	processedChannels := make([]chan DataPoint, numStreams)

	// Start data streams
	for i := 0; i < numStreams; i++ {
		dataChannels[i] = make(chan DataPoint, 100) // Buffered channel for each stream
		processedChannels[i] = make(chan DataPoint, 100)
		go generateData(ctx, dataChannels[i], i, time.Millisecond*100)
		go processData(dataChannels[i], processedChannels[i])
	}

	// Visualize processed data from all streams
	var wg sync.WaitGroup
	for i := 0; i < numStreams; i++ {
		wg.Add(1)
		go func(i int) {
			visualizeData(processedChannels[i])
			wg.Done()
		}(i)
	}

	// Wait for all data streams to finish processing
	wg.Wait()
}
