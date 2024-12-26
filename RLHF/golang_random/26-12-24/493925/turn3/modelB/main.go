package main

import (
	"fmt"
	"math/rand"
	"time"
)

func producer(id int, c chan<- int) {
	for {
		time.Sleep(time.Duration(rand.Intn(2000)) * time.Millisecond)
		c <- id
	}
}

func main() {
	channels := []chan int{}

	// Generate some random channels
	for i := 0; i < 5; i++ {
		c := make(chan int)
		channels = append(channels, c)
		go producer(i, c)
	}

	fmt.Println("Waiting for data from channels...")
	for {
		select {
		case data := <-channels[0]:
			fmt.Println("Received data from channel 0:", data)
		case data := <-channels[1]:
			fmt.Println("Received data from channel 1:", data)
		case data := <-channels[2]:
			fmt.Println("Received data from channel 2:", data)
		case data := <-channels[3]:
			fmt.Println("Received data from channel 3:", data)
		case data := <-channels[4]:
			fmt.Println("Received data from channel 4:", data)
		default:
			time.Sleep(time.Second)
			fmt.Println("No data available, checking again...")
		}
	}
}
