package main

import (
	"fmt"
	"sync"
)

// functional programming patterns
// map
func Map(slice []int, f func(int) int) []int {
	c := make(chan int)
	wg := sync.WaitGroup{}

	for _, v := range slice {
		wg.Add(1)
		go func(v int) {
			defer wg.Done()
			c <- f(v)
		}(v)
	}

	go func() {
		wg.Wait()
		close(c)
	}()

	result := []int{}
	for v := range c {
		result = append(result, v)
	}

	return result
}

// filter
func Filter(slice []int, f func(int) bool) []int {
	var result []int
	for _, v := range slice {
		if f(v) {
			result = append(result, v)
		}
	}
	return result
}

// reduce
func Reduce(slice []int, f func(int, int) int, initial int) int {
	result := initial
	for _, v := range slice {
		result = f(result, v)
	}
	return result
}

// generate a large slice
func generateLargeSlice(size int) []int {
	slice := make([]int, size)
	for i := 0; i < size; i++ {
		slice[i] = i
	}
	return slice
}

func main() {
	// generate a large slice for testing
	largeSlice := generateLargeSlice(1000000)

	// map operation
	mappedSlice := Map(largeSlice, func(x int) int {
		return x * 2
	})
	fmt.Println("Length of mapped slice:", len(mappedSlice))

	// filter operation
	filteredSlice := Filter(largeSlice, func(x int) bool {
		return x%2 == 0
	})
	fmt.Println("Length of filtered slice:", len(filteredSlice))

	// reduce operation
	sum := Reduce(largeSlice, func(x, y int) int {
		return x + y
	}, 0)
	fmt.Println("Sum of all elements in the slice:", sum)
}
