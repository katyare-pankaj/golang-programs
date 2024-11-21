package main

import (
	"fmt"
	"sync"
)

func produceNumbers(wg *sync.WaitGroup, ch chan int) {
	for i := 1; i <= 5; i++ {
		ch <- i
		fmt.Println("Produced:", i)
	}
	close(ch)
	wg.Done()
}

func printNumbers(wg *sync.WaitGroup, ch chan int) {
	wg.Wait() // Wait for the producer to finish
	for num := range ch {
		fmt.Println("Consumed:", num)
	}
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	ch := make(chan int)

	wg.Add(2) // We have two goroutines to wait for
	go produceNumbers(&wg, ch)
	go printNumbers(&wg, ch)

	wg.Wait() // Wait for all goroutines to finish
	fmt.Println("All done!")
}
