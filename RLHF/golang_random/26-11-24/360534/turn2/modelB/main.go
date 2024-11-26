package main

import (
	"fmt"
	"sync"
)

func generateData(dataCh chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 100; i++ {
		dataCh <- i
	}
	close(dataCh)
}

func transformData(dataCh <-chan int, transformedCh chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for data := range dataCh {
		transformedCh <- data + 1
	}
	close(transformedCh)
}

func aggregateResults(transformedCh <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	var total int
	for transformed := range transformedCh {
		total += transformed
	}
	fmt.Println("Final Sum:", total)
}

func main() {
	var wg sync.WaitGroup
	dataCh := make(chan int)
	transformedCh := make(chan int)

	// Launch the data generators
	wg.Add(1)
	go generateData(dataCh, &wg)

	// Launch the data transformers
	const numTransformers = 4
	wg.Add(numTransformers)
	for i := 0; i < numTransformers; i++ {
		go transformData(dataCh, transformedCh, &wg)
	}

	// Launch the aggregator
	wg.Add(1)
	go aggregateResults(transformedCh, &wg)

	wg.Wait()
}
