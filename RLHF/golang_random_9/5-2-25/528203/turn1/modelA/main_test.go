package main

import (
	"fmt"
	"testing"
	"time"
)

func inefficientFunction(data []int) int {
	sum := 0
	for _, value := range data {
		// This delay is just to simulate a heavier calculation
		time.Sleep(10 * time.Millisecond)
		sum += value
	}
	return sum
}

func main() {
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("Sum:", inefficientFunction(data))
}

func BenchmarkInefficientFunction(b *testing.B) {
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i := 0; i < b.N; i++ {
		inefficientFunction(data)
	}
}
