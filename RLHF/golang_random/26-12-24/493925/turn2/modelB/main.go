package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		ch1 <- 42
	}()

	go func() {
		time.Sleep(1 * time.Second)
		ch2 <- "hello"
	}()

	select {
	case msg := <-ch1:
		fmt.Println("Received int:", msg)
	case msg := <-ch2:
		fmt.Println("Received string:", msg)
	}
}
