package main

import (
	"fmt"
	"sync"
)

var sharedChannel = make(chan int)

func sendData(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		sharedChannel <- i
	}
	close(sharedChannel)
}
func receiveData(wg *sync.WaitGroup) {
	defer wg.Done()
	for data := range sharedChannel {
		fmt.Println("Received:", data)
	}
}
func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go sendData(&wg)
	go receiveData(&wg)
	wg.Wait()
}
