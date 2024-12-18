package main

import (
	"fmt"
	"time"
)

func producer(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
		fmt.Println("Produced:", i)
		time.Sleep(time.Second * 1)
	}
	close(ch)
}

func consumer(ch chan int) {
	for val := range ch {
		fmt.Println("Consumed:", val)
	}
}

func main() {
	ch := make(chan int, 3)
	go producer(ch)
	go consumer(ch)
	time.Sleep(time.Second * 15)
}
