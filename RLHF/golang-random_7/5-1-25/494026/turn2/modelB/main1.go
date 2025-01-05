package main

import (
	"fmt"
	"math"
	"os"
	"runtime"
	"runtime/pprof"
	"sync"
)

const (
	numTasks   = 100
	taskSize   = 100000
	numCPUs    = 4 // Set this to your number of logical CPUs
	maxRandNum = 1000000
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

func countPrimes(start, end int, primes chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	count := 0
	for i := start; i <= end; i++ {
		if isPrime(i) {
			count++
			primes <- i // Send the prime number to the channel
		}
	}
}

func main() {
	// Profiling memory allocations
	f, err := os.Create("mem.prof")
	if err != nil {
		fmt.Printf("Could not create memory profile file: %s\n", err)
	}
	defer f.Close()
	pprof.WriteHeapProfile(f)

	// Profiling CPU usage
	cpuFile, err := os.Create("cpu.prof")
	if err != nil {
		fmt.Printf("Could not create CPU profile file: %s\n", err)
	}
	defer cpuFile.Close()
	if err := pprof.StartCPUProfile(cpuFile); err != nil {
		fmt.Printf("Could not start CPU profile: %s\n", err)
	}
	defer pprof.StopCPUProfile()

	runtime.GOMAXPROCS(numCPUs)

	primes := make(chan int, taskSize*numTasks) // Channel to collect prime numbers
	var wg sync.WaitGroup

	// Distribute the work evenly among the tasks
	for i := 0; i < numTasks; i++ {
		start := i*taskSize + 1
		end := (i + 1) * taskSize
		wg.Add(1)
		go countPrimes(start, end, primes, &wg)
	}

	go func() {
		// Wait for all tasks to complete and close the channel
		wg.Wait()
		close(primes)
	}()

	// Consume the prime numbers from the channel
	var primeCount int
	for range primes { // Remove unused variable "prime"
		primeCount++
	}

	fmt.Println("Total primes found: ", primeCount)

	// Profiling heap usage after tasks complete
	pprof.WriteHeapProfile(f)

	// Analyze the CPU profile using "go tool pprof cpu.prof"
	// Analyze the memory profile using "go tool pprof mem.prof"
}
