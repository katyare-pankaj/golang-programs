package main

import (
	"fmt"
	"math"
	"sync"
)

func factorizeNumberConcurrently(n int) []int {
	var factors []int
	maxFactor := int(math.Sqrt(float64(n)))
	var wg sync.WaitGroup
	var found int32
	result := make(chan int)

	for i := 2; i <= maxFactor; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			for n%i == 0 {
				found++
				result <- i
				n /= i
			}
		}(i)
	}

	if n > 1 {
		factors = append(factors, n)
	}

	wg.Wait()

	for found > 0 {
		factors = append(factors, <-result)
		found--
	}

	return factors
}

// Main function to demonstrate factorizeNumberConcurrently.
func main() {
	number := 9999991
	fmt.Println("Factors of", number, ":", factorizeNumberConcurrently(number))
}
