package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type DataPoint struct {
	Timestamp time.Time
	Value     float64
}

func generateDataStream(id int, wg *sync.WaitGroup, dataCh chan<- DataPoint) {
	defer wg.Done()
	for {
		dataPoint := DataPoint{
			Timestamp: time.Now(),
			Value:     rand.Float64() * 100,
		}
		dataCh <- dataPoint
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	var wg sync.WaitGroup
	dataCh := make(chan DataPoint, 100)

	// Start multiple data streams
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go generateDataStream(i, &wg, dataCh)
	}

	// Update visualization in real-time
	go func() {
		for {
			select {
			case dataPoint := <-dataCh:
				fmt.Printf("Received data: %+v\n", dataPoint)
				// Update visualization here
			}
		}
	}()

	// Wait for all data streams to finish
	wg.Wait()
}
