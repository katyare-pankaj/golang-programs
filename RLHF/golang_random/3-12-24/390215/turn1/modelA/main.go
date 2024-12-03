package main

import (
	"fmt"
	"sync"
)

var sharedCounter int
var wg sync.WaitGroup
var mtx sync.Mutex

func incrementCounter() {
	for i := 0; i < 1000; i++ {
		mtx.Lock()
		sharedCounter++
		mtx.Unlock()
	}
	wg.Done()
}

func main() {
	wg.Add(5)

	for i := 0; i < 5; i++ {
		go incrementCounter()
	}

	wg.Wait()
	fmt.Println("Final counter value:", sharedCounter)
}
