package main

import (
	"fmt"
	"sync"
)

var sharedCounter int
var mutex sync.Mutex

func incrementCounter(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 10000; i++ {
		mutex.Lock()
		sharedCounter++
		mutex.Unlock()
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
