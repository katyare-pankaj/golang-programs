package main

import (
	"fmt"
	"time"
)

func createCachedFunction(f func(int) int, cacheSize int) func(int) int {
	cache := make(map[int]int)
	nextKey := 0

	return func(n int) int {
		if result, exists := cache[n]; exists {
			return result
		}

		result := f(n)
		cache[n] = result

		if len(cache) > cacheSize {
			delete(cache, nextKey)
			nextKey++
		}

		return result
	}
}

func expensiveFunction(n int) int {
	fmt.Printf("Computing %d^2...\n", n)
	time.Sleep(1 * time.Second)
	return n * n
}

func main() {
	cachedFunc := createCachedFunction(expensiveFunction, 3) // Cache size of 3

	fmt.Println(cachedFunc(5)) // Output: Computing 5^2... 25
	fmt.Println(cachedFunc(3)) // Output: Computing 3^2... 9
	fmt.Println(cachedFunc(5)) // Output: 25 (from cache)
	fmt.Println(cachedFunc(4)) // Output: Computing 4^2... 16
	fmt.Println(cachedFunc(2)) // Output: Computing 2^2... 4
	fmt.Println(cachedFunc(4)) // Output: 16 (from cache)
	fmt.Println(cachedFunc(6)) // Output: Computing 6^2... 36
	fmt.Println(cachedFunc(5)) // Output: Computing 5^2... 25
}
