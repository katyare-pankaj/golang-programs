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

type semaphore struct {
	m sync.Mutex
	v int
}

func (s *semaphore) Acquire(n int) {
	s.m.Lock()
	defer s.m.Unlock()
	for s.v < n {
		select {
		case <-time.After(time.Millisecond * 1):
		}
	}
	s.v -= n
}

func (s *semaphore) Release(n int) {
	s.m.Lock()
	defer s.m.Unlock()
	s.v += n
}

func generateData(ctx context.Context, id int, dataCh chan<- DataPoint) {
	defer close(dataCh)
	for {
		select {
		case <-ctx.Done():
			return
		default:
			dataCh <- DataPoint{
				Timestamp: time.Now(),
				Value:     rand.Float64() * 100,
			}
			time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		}
	}
}

func processData(ctx context.Context, dataCh chan DataPoint, sem *semaphore) {
	defer sem.Release(1)

	for {
		select {
		case <-ctx.Done():
			return
		case dataPoint, ok := <-dataCh:
			if !ok {
				return
			}
			fmt.Printf("Processed data: %+v\n", dataPoint)
		}
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	numStreams := 3
	sem := &semaphore{v: numStreams}

	dataCh := make(chan DataPoint, numStreams*10)
	processCh := make(chan DataPoint, numStreams*10)

	ctx, cancel := context.WithCancel(context.Background())

	// Start data generators
	for i := 0; i < numStreams; i++ {
		go generateData(ctx, i, dataCh)
	}

	// Start data processors
	for i := 0; i < numStreams; i++ {
		sem.Acquire(1)
		go processData(ctx, processCh, sem)
	}

	// Convert data stream to process stream
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case dataPoint, ok := <-dataCh:
				if !ok {
					return
				}
				processCh <- dataPoint
			}
		}
	}()

	select {}

	cancel()
}
