package main

import (
	"fmt"
	"os"
	"runtime"
	"runtime/pprof"
	"sync"
)

// sieve of Eratosthenes function
func sieve(start, end int, primes chan<- int) {
	isPrime := make([]bool, end-start+1)
	for i := 0; i < len(isPrime); i++ {
		isPrime[i] = true
	}

	// Process each prime number up to sqrt(end)
	for p := 2; p*p <= end; p++ {
		if isPrime[p-start] {
			for i := p*p - start; i < len(isPrime); i += p {
				isPrime[i] = false
			}
		}
	}

	// Collect primes
	for i := 2; i < len(isPrime); i++ {
		if isPrime[i] {
			primes <- start + i
		}
	}
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

	runtime.GOMAXPROCS(runtime.NumCPU())

	// Total range of numbers to check
	limit := 10000000
	primes := make(chan int, limit/10) // Use a buffer size based on expected number of primes
	var wg sync.WaitGroup

	// Divide the work into chunks
	chunkSize := limit / runtime.NumCPU()

	// Launch worker goroutines
	for i := 0; i < runtime.NumCPU(); i++ {
		start := i * chunkSize
		end := min(start+chunkSize-1, limit)
		wg.Add(1)
		go sieve(start, end, primes)
	}

	// Collect results
	totalPrimes := 0
	for prime := range primes {
		totalPrimes++
	}
	wg.Wait()
	close(primes)

	fmt.Printf("Total primes found: %d\n", totalPrimes)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
