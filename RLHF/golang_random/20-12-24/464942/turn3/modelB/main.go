package main

import (
	"fmt"
	"sync"
)

const numIncrements = 100000

func main() {
	var counter int
	var mutex sync.Mutex
	var wg sync.WaitGroup
	wg.Add(numIncrements)
	for i := 0; i < numIncrements; i++ {
		go incrementCounter(&counter, &mutex, &wg)
	}
	wg.Wait()
	fmt.Println("Final counter value:", counter)
}

func incrementCounter(counter *int, mutex *sync.Mutex, wg *sync.WaitGroup) {
	defer wg.Done()
	mutex.Lock()
	*counter++
	mutex.Unlock()
}
