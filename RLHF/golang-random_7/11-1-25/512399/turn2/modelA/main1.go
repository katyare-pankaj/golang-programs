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

type dataStream struct {
	id       int
	ctx      context.Context
	cancel   context.CancelFunc
	dataCh   <-chan DataPoint
	updateCh chan DataPoint
	rate     time.Duration
}

func createDataStream(id int, dataCh <-chan DataPoint, updateCh chan DataPoint, rate time.Duration) *dataStream {
	ctx, cancel := context.WithCancel(context.Background())
	return &dataStream{
		id:       id,
		ctx:      ctx,
		cancel:   cancel,
		dataCh:   dataCh,
		updateCh: updateCh,
		rate:     rate,
	}
}

func (ds *dataStream) start() {
	go func() {
		defer ds.cancel()

		for {
			select {
			case <-ds.ctx.Done():
				return
			case dataPoint, ok := <-ds.dataCh:
				if !ok {
					return
				}

				select {
				case ds.updateCh <- dataPoint:
					time.Sleep(ds.rate)
				default:
				}
			}
		}
	}()
}

func generateData(ctx context.Context, dataCh chan DataPoint) {
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

func updateVisualization(updateCh chan DataPoint) {
	var wg sync.WaitGroup
	numStreams := 3
	updateCh = make(chan DataPoint, numStreams*10)

	// Start data streams
	dataChannels := make([]chan DataPoint, numStreams) // change <-chan DataPoint to chan DataPoint
	for i := 0; i < numStreams; i++ {
		dataChannels[i] = make(chan DataPoint, 100)
		go generateData(context.Background(), dataChannels[i])
		wg.Add(1)
		stream := createDataStream(i, dataChannels[i], updateCh, time.Millisecond*100)
		stream.start()
	}

	// Simulate visualization update logic
	for {
		select {
		case dataPoint := <-updateCh:
			fmt.Printf("Received data: %+v\n", dataPoint)
			// Update visualization here
		}
	}

	wg.Wait()
}

func main() {
	rand.Seed(time.Now().UnixNano())

	// Start visualization update logic
	go updateVisualization(make(chan DataPoint, 100))

	select {}
}
