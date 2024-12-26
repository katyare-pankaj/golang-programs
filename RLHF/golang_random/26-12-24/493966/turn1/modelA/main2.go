package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(time.Duration(id) * time.Second)
	fmt.Printf("Worker %d finished\n", id)
}

func main() {
	var wg1, wg2 sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg1.Add(1)
		go worker(i, &wg1)
	}

	for i := 4; i <= 5; i++ {
		wg2.Add(1)
		go worker(i, &wg2)
	}

	wg1.Wait()
	fmt.Println("First group of workers has finished")

	wg2.Wait()
	fmt.Println("All workers have finished")
}
