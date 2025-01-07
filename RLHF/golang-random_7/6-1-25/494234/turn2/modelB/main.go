package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, slice []int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := id; i < len(slice); i += 2 {
		slice[i] = slice[i] * 2
		fmt.Printf("Worker %d: Updated slice[%d] = %d\n", id, i, slice[i])
		time.Sleep(time.Millisecond * 100)
	}
}

func main() {
	// Shared slice between workers
	sharedSlice := []int{1, 2, 3, 4, 5}

	var wg sync.WaitGroup
	wg.Add(2)

	go worker(0, sharedSlice, &wg)
	go worker(1, sharedSlice, &wg)

	wg.Wait()

	fmt.Println("Final slice state:", sharedSlice)
}
