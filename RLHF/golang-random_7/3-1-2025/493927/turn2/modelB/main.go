package main

import (
	"fmt"
	"sync"
	"time"
)

func processTask(taskId int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Processing task %d...\n", taskId)
	// Simulate task processing time
	time.Sleep(2 * time.Second)
}

func main() {
	var wg sync.WaitGroup
	tasks := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, taskId := range tasks {
		wg.Add(1)
		go processTask(taskId, &wg)
	}

	wg.Wait()
	fmt.Println("All tasks completed.")
}
