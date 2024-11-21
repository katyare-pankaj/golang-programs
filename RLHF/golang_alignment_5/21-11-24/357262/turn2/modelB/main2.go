package main

import (
	"fmt"
	"sync"
	"time"
)

func produceNumbers(wg *sync.WaitGroup, ch chan int) {
	for i := 1; i <= 10; i++ {
		ch <- i
		fmt.Println("Produced:", i)
		time.Sleep(time.Millisecond * 500)
	}
	close(ch)
	wg.Done()
}

func printNumbers(wg *sync.WaitGroup, ch chan int) {
	wg.Wait()
	for num := range ch {
		fmt.Println("Consumed:", num)
	}
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	ch := make(chan int, 3) // Buffer size of 3
	wg.Add(2)
	go produceNumbers(&wg, ch)
	go printNumbers(&wg, ch)
	wg.Wait()
	fmt.Println("All done!")
}
