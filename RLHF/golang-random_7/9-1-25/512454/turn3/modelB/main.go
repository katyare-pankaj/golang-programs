package main

import (
	"fmt"
	"sync"
)

func sumWorker(start int, end int, c chan int) {
	total := 0
	for i := start; i <= end; i++ {
		total += i
	}
	c <- total
}

func main() {
	var wg sync.WaitGroup
	c := make(chan int)
	n := 100000
	fmt.Printf("Calculating sum of integers from 1 to %d using concurrency:\n", n)

	numWorkers := 4 // You can adjust the number of workers as you like
	chunkSize := n/numWorkers + 1

	for i := 0; i < numWorkers; i++ {
		start := i*chunkSize + 1
		end := min((i+1)*chunkSize, n)
		wg.Add(1)
		go sumWorker(start, end, c)
	}

	go func() {
		wg.Wait()
		close(c)
	}()

	total := 0
	for sum := range c {
		total += sum
	}

	fmt.Println("Result:", total)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
