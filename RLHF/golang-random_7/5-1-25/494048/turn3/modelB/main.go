package main

import (
	"fmt"
)

func main() {

	arr := [...]int{1, 2, 3, 4, 5}

	for index, value := range arr {
		fmt.Println("Arry :", index, value)
	}

	slice := []string{"apple", "banana", "cherry"}

	for _, fruit := range slice {
		fmt.Println("Slice :", fruit)
	}

	maparr := map[string]int{"a": 1, "b": 2, "c": 3}

	for key, value := range maparr {
		fmt.Println("Map :", key, value)
	}

	// Iterate over a channel
	ch := make(chan int)
	go func() {
		ch <- 1
		ch <- 2
		close(ch)
	}()
	for value := range ch {
		fmt.Println("Channals :", value)
	}
}
