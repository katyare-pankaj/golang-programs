package main

import (
	"fmt"
	"time"
)

func main() {
	var mySlice []int
	for i := 0; i < 1000000; i++ {
		mySlice = append(mySlice, i)
	}
	time.Sleep(10 * time.Second)
	fmt.Println("Finished")
}
