package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(wg *sync.WaitGroup, id int, ch chan bool) {
	defer wg.Done()
	time.Sleep(time.Duration(id) * time.Second)
	fmt.Printf("Worker %d is done.\n", id)
	ch <- true
}

func main() {
	var wg sync.WaitGroup
	ch := make(chan bool, 10)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go worker(&wg, i, ch)
	}
	// Using wg.Wait() for fairness
	wg.Wait()
	// Using ch to check for starvation
	for i := 0; i < 10; i++ {
		select {
		case <-ch:
			fmt.Printf("Worker %d was done.\n", i)
		default:
			fmt.Printf("Worker %d is still running.\n", i)
		}
	}
}
