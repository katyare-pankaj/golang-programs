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

type DataStream struct {
	ID   string
	Data chan DataPoint
	wg   sync.WaitGroup
	done chan struct{}
}

func (ds *DataStream) Start() {
	ds.wg.Add(1)
	go func() {
		defer ds.wg.Done()
		for {
			select {
			case <-ds.done:
				return
			default:
				// Generate random data points for demonstration purposes
				dataPoint := DataPoint{
					Timestamp: time.Now(),
					Value:     rand.Float64(),
				}
				ds.Data <- dataPoint
				time.Sleep(time.Second * 1) // Simulate data generation interval
			}
		}
	}()
}

func (ds *DataStream) Stop() {
	close(ds.done)
	ds.wg.Wait()
	close(ds.Data)
}

func main() {
	// Create multiple data streams
	streams := []*DataStream{
		{ID: "stream1", Data: make(chan DataPoint, 100)},
		{ID: "stream2", Data: make(chan DataPoint, 100)},
		{ID: "stream3", Data: make(chan DataPoint, 100)},
	}

	// Start all data streams
	for _, stream := range streams {
		stream.Start()
	}

	// Visualize data from all streams concurrently
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-time.After(time.Second * 5): // Refresh plot every 5 seconds
				fmt.Println("\nRefreshing plot:")
				for _, stream := range streams {
					fmt.Printf("Stream %s: \n", stream.ID)
					for dataPoint := range stream.Data {
						fmt.Printf("\t%s: %f\n", dataPoint.Timestamp.Format("15:04:05"), dataPoint.Value)
					}
				}
			}
		}
	}()

	// Simulate data streaming for a period
	time.Sleep(time.Second * 30)

	// Stop all data streams
	for _, stream := range streams {
		stream.Stop()
	}

	// Wait for visualization to finish
	wg.Wait()
}
