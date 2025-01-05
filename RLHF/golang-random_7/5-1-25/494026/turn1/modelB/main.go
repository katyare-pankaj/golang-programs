package main

import (
	"fmt"
	"math"
	"os"
	"runtime"
	"runtime/pprof"
	"sync"
)

func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func countPrimes(start, end int, wg *sync.WaitGroup) {
	defer wg.Done()
	count := 0
	for i := start; i <= end; i++ {
		if isPrime(i) {
			count++
		}
	}
	fmt.Println("Primes found: ", count)
}

func main() {
	// Profiling memory allocations
	f, err := os.Create("mem.prof")
	if err != nil {
		fmt.Printf("Could not create memory profile file: %s\n", err)
	}
	defer f.Close()
	pprof.WriteHeapProfile(f)

	// Create a WaitGroup to synchronize concurrent tasks
	var wg sync.WaitGroup

	// Let's assume we have 8 logical cores on our system
	numCPUs := runtime.NumCPU()
	runtime.GOMAXPROCS(numCPUs)

	const numTasks = 100   // number of tasks to divide the work into
	const taskSize = 10000 // size of each task

	fmt.Println("Number of CPUs: ", numCPUs)

	// Distribute the work evenly among the tasks
	for i := 0; i < numTasks; i++ {
		start := i*taskSize + 1
		end := (i + 1) * taskSize
		wg.Add(1)
		go countPrimes(start, end, &wg)
	}

	// Wait for all tasks to complete
	wg.Wait()

	// Profiling CPU usage
	f, err = os.Create("cpu.prof")
	if err != nil {
		fmt.Printf("Could not create CPU profile file: %s\n", err)
	}
	defer f.Close()
	pprof.Lookup("goroutine").WriteTo(f, 2)
}
