package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	// Starting several goroutines
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			fmt.Printf("Goroutine %d starting\n", id)
			time.Sleep(time.Duration(id) * time.Second)
			fmt.Printf("Goroutine %d done\n", id)
		}(i)
	}

	// Waiting for all goroutines to finish
	wg.Wait()
	fmt.Println("All goroutines finished")
}
