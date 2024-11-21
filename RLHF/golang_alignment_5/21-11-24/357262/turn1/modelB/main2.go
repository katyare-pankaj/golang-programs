package main

import (
	"fmt"
	"time"
)

func producer(ch chan int) {
	for i := 1; i <= 10; i++ {
		ch <- i
		fmt.Println("Sent:", i)
		time.Sleep(1 * time.Second)
	}
	close(ch)
}

func consumer(ch chan int) {
	for {
		num := <-ch
		fmt.Println("Received:", num)
	}
}

func main() {
	ch := make(chan int) // Unbuffered channel
	go producer(ch)
	go consumer(ch)
	time.Sleep(12 * time.Second)
}
