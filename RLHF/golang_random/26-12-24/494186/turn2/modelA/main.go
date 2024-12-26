package main

import (
	"fmt"
	"sync"
	"time"
)

func producer(id int, ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 5; i++ {
		num := id*10 + i
		ch <- num
		fmt.Println("Producer", id, "sent:", num)
		time.Sleep(time.Second / 2)
	}
}

func consumer(id int, ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for num := range ch {
		fmt.Println("Consumer", id, "received:", num)
		time.Sleep(time.Second / 2)
	}
}

func main() {
	var wg sync.WaitGroup
	ch := make(chan int)

	wg.Add(2) // 2 producers
	go producer(1, ch, &wg)
	go producer(2, ch, &wg)

	wg.Add(1) // 1 consumer
	go consumer(3, ch, &wg)

	wg.Wait()
	close(ch) // Inform all consumers that channel is closed
}
