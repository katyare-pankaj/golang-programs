package main

import (
	"fmt"
	"sync"
)

// Example function that might panic
func potentialPanic() {
	// This will panic if i == 0
	i := 0
	fmt.Printf("1/%d\n", i)
}

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()

	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Worker %d recovered from panic: %v\n", id, r)
		}
	}()

	potentialPanic()
	fmt.Printf("Worker %d finished executing\n", id)
}

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}

	wg.Wait()
	fmt.Println("Main function finished")
}
