package main

import (
	"fmt"
	"sync"
)

func producer(mutex *sync.Mutex, cond *sync.Cond, data *[]int) {
	for i := 0; i < 5; i++ {
		mutex.Lock()
		*data = append(*data, i)
		fmt.Println("Produced:", i)
		cond.Signal()
		mutex.Unlock()
	}
}

func consumer(mutex *sync.Mutex, cond *sync.Cond, data *[]int) {
	for i := 0; i < 5; i++ {
		mutex.Lock()
		for len(*data) == 0 {
			cond.Wait()
		}
		item := (*data)[0]
		*data = (*data)[1:]
		fmt.Println("Consumed:", item)
		mutex.Unlock()
	}
}

func main() {
	var data []int
	mutex := &sync.Mutex{}
	cond := sync.NewCond(mutex)

	go producer(mutex, cond, &data)
	go consumer(mutex, cond, &data)

	// Wait for goroutines to finish
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		wg.Wait()
	}()
}
