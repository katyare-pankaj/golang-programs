package main

import (
	"flag"
	"fmt"
	"sync"
	"time"
)

var (
	delay1 time.Duration
	delay2 time.Duration
	delay3 time.Duration
)

func init() {
	// Define flags for command-line arguments
	flag.DurationVar(&delay1, "delay1", 2*time.Second, "Delay time for the first task (default is 2s)")
	flag.DurationVar(&delay2, "delay2", 1*time.Second, "Delay time for the second task (default is 1s)")
	flag.DurationVar(&delay3, "delay3", 3*time.Second, "Delay time for the third task (default is 3s)")
}

func task(wg *sync.WaitGroup, id int, delay time.Duration) {
	defer wg.Done()
	fmt.Printf("Task %d started with a delay of %s\n", id, delay)
	time.Sleep(delay)
	fmt.Printf("Task %d completed.\n", id)
}

func main() {
	// Parse command-line flags
	flag.Parse()

	var wg sync.WaitGroup

	// Start the three tasks with user-defined delays
	wg.Add(1)
	go task(&wg, 1, delay1)

	wg.Add(1)
	go task(&wg, 2, delay2)

	wg.Add(1)
	go task(&wg, 3, delay3)

	// Wait for all tasks to finish
	wg.Wait()

	fmt.Println("All tasks completed. Main function exiting...")
}
