package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func task(taskNo int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Starting task %d...\n", taskNo)
	sleepDuration := time.Duration(rand.Intn(3)+1) * time.Second //random sleep from 1 to 3 seconds
	time.Sleep(sleepDuration)
	fmt.Printf("Task %d completed after sleeping for %s\n", taskNo, sleepDuration)
}

func main() {
	var wg sync.WaitGroup
	fmt.Println("Starting tasks concurrently...")

	numTasks := 3
	wg.Add(numTasks)
	//Run each task in a separate goroutine
	for i := 1; i <= numTasks; i++ {
		go task(i, &wg)
	}

	wg.Wait() //Wait for all tasks to finish
	fmt.Println("All tasks completed.")
}
