package main

import (
	"fmt"
	"math"
	"runtime"
	"sync"
	"time"
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

func calculatePrimes(limit int) []int {
	var primes []int
	for i := 2; i <= limit; i++ {
		if isPrime(i) {
			primes = append(primes, i)
		}
	}
	return primes
}

func sieveOfEratosthenes(limit int) []int {
	primes := make([]int, 0)
	isPrime := make([]bool, limit+1)
	for i := 2; i <= limit; i++ {
		isPrime[i] = true
	}

	for i := 2; i*i <= limit; i++ {
		if isPrime[i] {
			for j := i * i; j <= limit; j += i {
				isPrime[j] = false
			}
		}
	}

	for i := 2; i <= limit; i++ {
		if isPrime[i] {
			primes = append(primes, i)
		}
	}
	return primes
}

func calculatePrimesConcurrently(limit int) []int {
	var primes []int
	var wg sync.WaitGroup

	numThreads := runtime.NumCPU()
	chunkSize := (limit-1)/numThreads + 1

	wg.Add(numThreads)
	for threadID := 0; threadID < numThreads; threadID++ {
		start := 2 + threadID*chunkSize
		end := min(start+chunkSize-1, limit)
		go func(start, end int) {
			defer wg.Done()
			for i := start; i <= end; i++ {
				if isPrime(i) {
					primes = append(primes, i)
				}
			}
		}(start, end)
	}

	wg.Wait()
	return primes
}

func main() {
	limit := 100000
	start := time.Now()
	primes := calculatePrimes(limit)
	duration := time.Since(start)
	fmt.Printf("Calculated %d primes up to %d in %s\n", len(primes), limit, duration)

	start = time.Now()
	primes = sieveOfEratosthenes(limit)
	duration = time.Since(start)
	fmt.Printf("Calculated %d primes using Sieve of Eratosthenes in: %s\n", len(primes), duration)

	start = time.Now()
	primes = calculatePrimesConcurrently(limit)
	duration = time.Since(start)
	fmt.Printf("Calculated %d primes concurrently in: %s\n", len(primes), duration)
}
