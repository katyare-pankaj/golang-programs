package main

import (
	"fmt"
	"time"
)

func worker(done chan struct{}, result chan int) {
	time.Sleep(time.Second)
	result <- 42
	close(done)
	fmt.Println("Worker done")
}

func main() {
	done := make(chan struct{})
	result := make(chan int)

	go worker(done, result)

	select {
	case <-done:
		val := <-result
		fmt.Println("Worker completed:", val)
	case <-time.After(time.Second * 2):
		fmt.Println("Worker timed out")
	}
}
