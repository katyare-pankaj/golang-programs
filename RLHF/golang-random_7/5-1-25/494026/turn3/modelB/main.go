package main

import (
	"fmt"
	"os"
	"runtime"
	"runtime/pprof"
)

const (
	numTasks = 4
	taskSize = 1000000
)

func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func countPrimes(start, end int, countChan chan<- int) {
	count := 0
	for i := start; i <= end; i++ {
		if isPrime(i) {
			count++
		}
	}
	countChan <- count
}

func main() {
	// Profiling CPU usage
	f, err := os.Create("cpu.prof")
	if err != nil {
		fmt.Printf("Could not create CPU profile file: %s\n", err)
	}
	defer f.Close()
	if err := pprof.StartCPUProfile(f); err != nil {
		fmt.Printf("Could not start CPU profile: %s\n", err)
	}
	defer pprof.StopCPUProfile()

	runtime.GOMAXPROCS(numTasks)
	countChan := make(chan int, numTasks)

	// Distribute the work
	for i := 0; i < numTasks; i++ {
		start := i*taskSize + 1
		end := (i + 1) * taskSize
		go countPrimes(start, end, countChan)
	}

	// Collect results
	totalCount := 0
	for i := 0; i < numTasks; i++ {
		count := <-countChan
		totalCount += count
	}

	fmt.Println("Total primes found: ", totalCount)
}
