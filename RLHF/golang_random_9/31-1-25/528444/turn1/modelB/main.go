package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Metric struct {
	CPU    float64
	Memory float64
}

var metrics []Metric
var mutex sync.Mutex

func updateMetrics(wg *sync.WaitGroup, updateChan chan<- Metric) {
	defer wg.Done()
	for {
		time.Sleep(time.Millisecond * 2)
		newMetric := Metric{
			CPU:    rand.Float64() * 100,
			Memory: rand.Float64() * 100,
		}
		// Send the new metric to the update channel
		updateChan <- newMetric
	}
}

func monitorMetrics(wg *sync.WaitGroup, updateChan <-chan Metric) {
	defer wg.Done()
	for {
		select {
		case newMetric := <-updateChan:
			mutex.Lock()
			metrics = append(metrics, newMetric)
			mutex.Unlock()

		case <-time.After(time.Millisecond * 5):
			mutex.Lock()
			fmt.Println("Current Metrics:")
			for _, metric := range metrics {
				fmt.Printf("CPU: %.2f%% Memory: %.2f%%\n", metric.CPU, metric.Memory)
			}
			mutex.Unlock()
		}
	}
}

func main() {
	wg := &sync.WaitGroup{}
	updateChan := make(chan Metric)

	wg.Add(2)
	go updateMetrics(wg, updateChan)
	go monitorMetrics(wg, updateChan)

	wg.Wait()
}
