package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
var counter int
var c = &sync.Cond{}

func incCounter() {
	c.L.Lock()
	defer c.L.Unlock()

	counter++
	fmt.Printf("Incremented counter: %d\n", counter)
	c.Broadcast()
}

func waitForCounter(target int) {
	defer wg.Done()

	c.L.Lock()
	for counter < target {
		c.Wait()
	}
	fmt.Printf("Counter reached %d\n", target)
	c.L.Unlock()
}

func main() {
	wg.Add(3)

	go incCounter()
	go incCounter()
	go incCounter()

	go waitForCounter(3)
	go waitForCounter(3)
	go waitForCounter(3)

	wg.Wait()
}
