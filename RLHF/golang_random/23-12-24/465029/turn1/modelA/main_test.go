package main

import (
	"math/rand"
	"sync"
	"testing"
	"time"
)

func TestFactorialBenchmark(t *testing.T) {
	t.Run("FactorialParallel", func(t *testing.T) {
		b := testing.B{}
		b.SetParallelism(testing.GOMAXPROCS(0))
		for n := 0; n < b.N; n++ {
			factorialParallel(rand.Intn(20))
		}
	})
}

func factorial(n int) int64 {
	if n == 0 || n == 1 {
		return 1
	}
	return int64(n) * factorial(n-1)
}

func factorialParallel(n int) {
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		result := factorial(n)
		defer wg.Done()
		// Do something with the result if needed
		// In a real application, the result might be accumulated or processed
	}()
	wg.Wait()
}

func init() {
	rand.Seed(time.Now().UnixNano())
}
