package main

import (
	"fmt"
	"math/rand"
	"time"
)

func sumAndCount(numbers []int) (sum int, count int) {
	sum = 0
	count = 0
	for _, num := range numbers {
		sum += num
		count++
	}
	return
}

func main() {
	// Generate a large slice of random numbers
	rand.Seed(time.Now().UnixNano())
	n := 1000000
	numbers := make([]int, n)
	for i := range numbers {
		numbers[i] = rand.Intn(1000)
	}

	// Measure performance of using multiple return values
	start := time.Now()
	for i := 0; i < 1000; i++ {
		_, _ = sumAndCount(numbers)
	}
	end := time.Now()
	fmt.Printf("Time taken with multiple return values: %v\n", end.Sub(start))

	// Measure performance of using a single struct return value
	start = time.Now()
	for i := 0; i < 1000; i++ {
		result := sumAndCountStruct(numbers)
		_ = result.Sum
		_ = result.Count
	}
	end := time.Now()
	fmt.Printf("Time taken with struct return value: %v\n", end.Sub(start))
}

type result struct {
	Sum   int
	Count int
}

func sumAndCountStruct(numbers []int) result {
	sum := 0
	count := 0
	for _, num := range numbers {
		sum += num
		count++
	}
	return result{sum, count}
}
