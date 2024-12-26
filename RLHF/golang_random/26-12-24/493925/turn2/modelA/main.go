package main

import (
	"fmt"
	"time"
)

func signalOne() chan bool {
	c := make(chan bool)
	go func() {
		time.Sleep(time.Second * 2)
		c <- true
	}()
	return c
}

func signalTwo() chan bool {
	c := make(chan bool)
	go func() {
		time.Sleep(time.Second * 1)
		c <- true
	}()
	return c
}

func main() {
	signal1 := signalOne()
	signal2 := signalTwo()

	fmt.Println("Waiting for signals...")
	select {
	case <-signal1:
		fmt.Println("Received signal from signalOne.")
	case <-signal2:
		fmt.Println("Received signal from signalTwo.")
	}
}
