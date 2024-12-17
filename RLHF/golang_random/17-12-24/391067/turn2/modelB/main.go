package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

// Function with panic recovery
func processTask(wg *sync.WaitGroup, id int, taskCh <-chan int) {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("Goroutine %d panicked: %v", id, r)
		}
		wg.Done()
	}()

	for num := range taskCh {
		if num%2 == 0 {
			panic(fmt.Sprintf("Goroutine %d panicked with input %d", id, num))
		}
		fmt.Printf("Goroutine %d processed %d\n", id, num)
	}
}

func main() {
	var wg sync.WaitGroup
	taskCh := make(chan int, 10)

	// Launch multiple goroutines
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go processTask(&wg, i, taskCh)
	}

	// Send tasks to the channel
	go func() {
		defer close(taskCh)
		for i := 1; i <= 10; i++ {
			taskCh <- i
			time.Sleep(100 * time.Millisecond)
		}
	}()

	// Wait for all goroutines to finish
	wg.Wait()
	fmt.Println("All goroutines completed.")
}
