package main

import (
	"fmt"
	"sync"
)

var sharedCounter int
var wg sync.WaitGroup
var done chan struct{}

func incrementCounter() {
	for i := 0; i < 1000; i++ {

		sharedCounter++
	}
	wg.Done()
	done <- struct{}{}
}

func main() {
	wg.Add(5)

	for i := 0; i < 5; i++ {
		go incrementCounter()
	}

	for i := 0; i < 5; i++ {
		<-done
	}

	fmt.Println("Final counter value:", sharedCounter)
}
