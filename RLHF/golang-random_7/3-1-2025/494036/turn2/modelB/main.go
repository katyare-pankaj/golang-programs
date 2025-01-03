package main

import (
	"fmt"
	"sync"
	"time"
)

func task(wg *sync.WaitGroup, id int, delay time.Duration) {
	defer wg.Done()
	fmt.Printf("Task %d started with a delay of %s\n", id, delay)
	time.Sleep(delay)
	fmt.Printf("Task %d completed.\n", id)
}

func main() {
	var wg sync.WaitGroup
	numTasks := 3
	delays := []time.Duration{2 * time.Second, 1 * time.Second, 3 * time.Second}

	for i := 0; i < numTasks; i++ {
		wg.Add(1)
		go task(&wg, i+1, delays[i])
	}
	wg.Wait()
	fmt.Println("All tasks completed. Main function exiting...")
}
