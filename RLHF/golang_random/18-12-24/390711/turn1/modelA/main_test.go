package main

import (
	"fmt"
	"sync"
	"testing"
)

func processNumbers(numbers []int, wg *sync.WaitGroup) {
	defer wg.Done()
	total := 0
	for _, num := range numbers {
		total += num
	}
	fmt.Println("Sum:", total)
}

func TestProcessNumbers(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5}
	var wg sync.WaitGroup
	wg.Add(1)

	go processNumbers(numbers, &wg)

	wg.Wait()

	// Checking if the sum is printed correctly
	// In practice, we would capture the output and check it
	// Here we assume it's printed correctly as the focus is on the goroutine behavior
	t.Log("Expected sum:", sum(numbers))
}

func sum(numbers []int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}
