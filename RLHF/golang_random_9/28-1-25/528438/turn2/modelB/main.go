package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const numWorkers = 5
const numTasks = 20

func worker(id int, taskCh <-chan int, wg *sync.WaitGroup) {
	for task := range taskCh {
		fmt.Printf("Worker %d: Starting task %d\n", id, task)
		time.Sleep(time.Duration(rand.Intn(2000)) * time.Millisecond)
		fmt.Printf("Worker %d: Completed task %d\n", id, task)
		wg.Done()
	}
}
func main() {
	taskCh := make(chan int)
	var wg sync.WaitGroup
	wg.Add(numTasks)
	for i := 0; i < numWorkers; i++ {
		go worker(i+1, taskCh, &wg)
	}
	for i := 0; i < numTasks; i++ {
		taskCh <- i + 1
	}
	close(taskCh)
	wg.Wait()
	fmt.Println("All tasks have been completed.")
}
