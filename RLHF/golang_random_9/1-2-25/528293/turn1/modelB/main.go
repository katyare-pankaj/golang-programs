package main

import (
	"fmt"
	"sync"
	"time"
)

func doWork(f func() int) int {
	return f()
}

func main() {
	// Closure example
	var wg sync.WaitGroup
	numWorkers := 10
	workCount := 100000
	wg.Add(numWorkers)

	start := time.Now()
	for i := 0; i < numWorkers; i++ {
		go func(workerId int) {
			for j := 0; j < workCount; j++ {
				_ = doWork(func() int {
					return workerId * j
				})
			}
			wg.Done()
		}(i)
	}
	wg.Wait()
	elapsed := time.Since(start)
	fmt.Printf("Closure example: Elapsed time: %s\n", elapsed)

	// Callback example with closure
	start = time.Now()
	for i := 0; i < numWorkers; i++ {
		go func(workerId int) {
			for j := 0; j < workCount; j++ {
				doWorkWithCallback(func(result int) {}, workerId*j)
			}
			wg.Done()
		}(i)
	}
	wg.Wait()
	elapsed = time.Since(start)
	fmt.Printf("Callback example with closure: Elapsed time: %s\n", elapsed)

	// Callback example with function value
	start = time.Now()
	for i := 0; i < numWorkers; i++ {
		go func(workerId int) {
			for j := 0; j < workCount; j++ {
				doWorkWithCallback(func(result int) {}, workerId*j)
			}
			wg.Done()
		}(i)
	}
	wg.Wait()
	elapsed = time.Since(start)
	fmt.Printf("Callback example with function value: Elapsed time: %s\n", elapsed)
}

func doWorkWithCallback(callback func(int), work int) {
	result := work
	callback(result)
}
