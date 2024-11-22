package main

import (
	"fmt"
	"sync"
	"time"

	"golang.org/x/exp/rand"
)

// define a Simple Data structure
type Task struct {
	id   int
	data string
}

func main() {
	// Create a buffered channel with a buffer size of 3
	taskCh := make(chan Task, 3)
	var wg sync.WaitGroup

	// Launch three producer goroutines
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go producer(taskCh, i, &wg)
	}

	// Launch a single consumer goroutine
	wg.Add(1)
	go consumer(taskCh, &wg)

	// Wait for all goroutines to finish
	wg.Wait()
	fmt.Println("All tasks completed!")
}

// Producer function generates data and sends it to the channel
func producer(ch chan<- Task, id int, wg *sync.WaitGroup) {
	defer wg.Done()
	for j := 1; j <= 5; j++ {
		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
		task := Task{id: id, data: fmt.Sprintf("Task %d.%d", id, j)}
		ch <- task // Send data to the channel
		fmt.Println("Producer", id, "produced:", task.data)
	}
	close(ch) // Close the channel when done producing
}

// Consumer function receives data from the channel and processes it
func consumer(ch <-chan Task, wg *sync.WaitGroup) {
	defer wg.Done()
	for task := range ch {
		fmt.Println("Consumer processed:", task.data)
		time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond)
	}
}
