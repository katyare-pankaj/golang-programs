package main

import (
	"fmt"
	"sync"
)

const chunkSize = 100

func calculatePartialSum(start, end int, partialSums chan int) {
	partialSum := 0
	for i := start; i <= end; i++ {
		partialSum += i
	}
	partialSums <- partialSum
}

func main() {
	var wg sync.WaitGroup
	partialSums := make(chan int)

	n := 10000 // The upper limit for the sum
	numChunks := (n / chunkSize) + 1

	fmt.Printf("Calculating sum of integers from 1 to %d using concurrency:\n", n)

	for chunk := 0; chunk < numChunks; chunk++ {
		start := chunk*chunkSize + 1
		end := min((chunk+1)*chunkSize, n)
		wg.Add(1)
		go func() {
			defer wg.Done()
			calculatePartialSum(start, end, partialSums)
		}()
	}

	var totalSum int
	go func() {
		wg.Wait()
		close(partialSums)
	}()

	for partial := range partialSums {
		totalSum += partial
	}

	fmt.Println("Total sum:", totalSum)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
