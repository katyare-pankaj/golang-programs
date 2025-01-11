package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	numbers := []int{1, 2, 3, 4, 5}

	for i, number := range numbers {
		wg.Add(1)
		go func(i int, number int) {
			fmt.Println("Goroutine: index =", i, ", number =", number)
			// You can use the captured variables here
			wg.Done()
		}(i, number)
	}

	wg.Wait()
	fmt.Println("All goroutines have finished.")
}
