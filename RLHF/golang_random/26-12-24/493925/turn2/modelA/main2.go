package main

import (
	"fmt"
	"time"
)

func main() {
	timeout := make(chan bool, 1)
	go func() {
		time.Sleep(time.Second * 3)
		timeout <- true
	}()

	fmt.Println("Waiting for signal or timeout...")
	select {
	case <-timeout:
		fmt.Println("Timed out.")
	default:
		fmt.Println("Did not time out.")
	}
}
