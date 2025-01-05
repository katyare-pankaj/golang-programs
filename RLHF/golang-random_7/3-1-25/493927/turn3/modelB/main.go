package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func processTask(taskId int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Processing task %d...\n", taskId)

	// Simulate random processing time between 1 and 5 seconds
	processingTime := time.Duration(rand.Intn(5)+1) * time.Second
	time.Sleep(processingTime)
}

func main() {
	rand.Seed(time.Now().UnixNano()) // Seed the random number generator
	var wg sync.WaitGroup
	tasks := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, taskId := range tasks {
		wg.Add(1)
		go processTask(taskId, &wg)
	}

	wg.Wait()
	fmt.Println("All tasks completed.")
}
