package main

import (
	"errors"
	"fmt"
	"time"
)

func worker(id int, done chan bool, errChan chan error) {
	defer close(errChan)
	defer close(done)
	time.Sleep(time.Second * 2)
	if id == 2 {
		errChan <- errors.New("worker 2 failed")
		return
	}
	fmt.Println("Worker", id, "completed")
}

func main() {
	numWorkers := 3
	done := make(chan bool)
	errChan := make(chan error)

	for i := 1; i <= numWorkers; i++ {
		go worker(i, done, errChan)
	}

	for i := 1; i <= numWorkers; i++ {
		select {
		case <-done:
			fmt.Println("Worker", i, "done")
		case err := <-errChan:
			fmt.Println("Error:", err)
		}
	}
}
