package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func worker(id int, dur time.Duration) {
	defer wg.Done()
	fmt.Printf("Worker %d starting...\n", id)
	time.Sleep(dur)
	fmt.Printf("Worker %d done!\n", id)
}

func main() {
	wg.Add(3)
	defer wg.Wait()

	go worker(1, time.Second*1)
	go worker(2, time.Second*2)
	go worker(3, time.Second*3)

	fmt.Println("Waiting for workers...")
}
