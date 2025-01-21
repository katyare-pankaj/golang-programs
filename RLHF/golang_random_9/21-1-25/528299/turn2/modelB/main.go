package main

import (
	"fmt"
	"sync"
)

func main() {
	mySlice := []int{1, 2, 3, 4, 5}
	var mutex sync.Mutex

	// Start a goroutine to modify the slice
	go func() {
		mutex.Lock()
		defer mutex.Unlock()
		mySlice = append(mySlice, 6, 7, 8)
	}()

	// Iterate over the slice using a for loop
	for {
		mutex.Lock()
		if len(mySlice) == 0 {
			mutex.Unlock()
			break
		}
		element := mySlice[0]
		mySlice = mySlice[1:]
		mutex.Unlock()

		fmt.Println("Element:", element)
	}

	fmt.Println("Slice is empty.")
}
