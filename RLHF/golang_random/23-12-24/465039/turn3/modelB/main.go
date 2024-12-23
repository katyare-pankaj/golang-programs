package main

import (
	"fmt"
	"sync"
	"time"
)

// ProcessNum simulates work for each number
func ProcessNum(num int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Processing number: %d\n", num)
	// Simulate work using time.Sleep
	time.Sleep(time.Duration(num) * time.Millisecond)
}

func main() {
	numWorkers := 4
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var wg sync.WaitGroup

	wg.Add(len(numbers))
	for _, num := range numbers {
		go ProcessNum(num, &wg)
	}

	wg.Wait()
	fmt.Println("All numbers processed.")
}
