package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"

	"github.com/wcharczuk/go-charts/v2"
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

func fanInDataStreams(wg *sync.WaitGroup, dataChannels ...<-chan DataPoint) <-chan DataPoint {
	var fanInCh = make(chan DataPoint)
	go func() {
		defer close(fanInCh)
		defer wg.Done()
		var wg sync.WaitGroup
		for _, ch := range dataChannels {
			wg.Add(1)
			go func(c <-chan DataPoint) {
				defer wg.Done()
				for dp := range c {
					fanInCh <- dp
				}
			}(ch)
		}
		wg.Wait()
	}()
	return fanInCh
}

func rateLimitDataStream(ch <-chan DataPoint, limit time.Duration) <-chan DataPoint {
	limitedCh := make(chan DataPoint)
	go func() {
		defer close(limitedCh)
		ticker := time.NewTicker(limit)
		defer ticker.Stop()
		for dp := range ch {
			select {
			case <-ticker.C:
				limitedCh <- dp
			}
		}
	}()
	return limitedCh
}

func updateVisualization(dataCh <-chan DataPoint) {
	lineChart := charts.NewLine()
	lineChart.Width = 600
	lineChart.Height = 400
	lineChart.Series = append(lineChart.Series, charts.NewLineSeries("Data Stream"))
	var mutex sync.Mutex

	for {
		select {
		case dataPoint := <-dataCh:
			mutex.Lock()
			lineChart.Series[0].Points = append(lineChart.Series[0].Points, charts.DataPoint{X: dataPoint.Timestamp, Y: dataPoint.Value})
			mutex.Unlock()
			// Save or render the chart here
		}
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	var wg sync.WaitGroup
	dataChannels := make([]<-chan DataPoint, 3)