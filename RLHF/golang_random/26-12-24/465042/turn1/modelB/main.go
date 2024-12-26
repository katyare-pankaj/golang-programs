package main

import (
	"fmt"
	"sync"
)

var sharedCounter int // Shared variable
func incrementCounter(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 10000; i++ {
		sharedCounter++
	}
}
func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go incrementCounter(&wg)
	go incrementCounter(&wg)
	wg.Wait()
	fmt.Println("Final value of sharedCounter:", sharedCounter)
}
