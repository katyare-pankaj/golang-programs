package main

import (
	"fmt"
	"sync"
	"time"
)

func processTask(taskId int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Starting task %d\n", taskId)
	time.Sleep(time.Duration(100+taskId) * time.Millisecond)
	fmt.Printf("Task %d completed\n", taskId)
}

func main() {
	var wg sync.WaitGroup
	numTasks := 5

	wg.Add(numTasks)
	for i := 1; i <= numTasks; i++ {
		go processTask(i, &wg)
	}

	// Wait for all tasks to complete
	wg.Wait()

	fmt.Println("All tasks completed. Exiting program.")
}
