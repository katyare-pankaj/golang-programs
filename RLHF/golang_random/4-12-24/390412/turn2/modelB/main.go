package main

import (
	"fmt"
	"runtime"
	"sync"
)

func squareSum(start, end int, wg *sync.WaitGroup, resultChan chan<- int) {
	defer wg.Done()
	sum := 0
	for i := start; i <= end; i++ {
		sum += i * i
	}
	resultChan <- sum
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	const n = 100000
	const numGoroutines = 10
	var wg sync.WaitGroup
	resultChan := make(chan int, numGoroutines)

	chunkSize := n / numGoroutines

	wg.Add(numGoroutines)
	for i := 0; i < numGoroutines; i++ {
		start := i*chunkSize + 1
		end := start + chunkSize - 1
		if i == numGoroutines-1 {
			end = n
		}
		go squareSum(start, end, &wg, resultChan)
	}

	wg.Wait()
	close(resultChan)

	totalSum := 0
	for result := range resultChan {
		totalSum += result
	}

	fmt.Printf("Total sum of squares: %d\n", totalSum)
}
