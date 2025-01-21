package main

import (
	"fmt"
)

func main() {
	//original slice
	mySlice := []int{1, 2, 3, 4, 5}

	//create a copy of the slice
	copySlice := make([]int, len(mySlice))
	copy(copySlice, mySlice)

	for index, value := range copySlice {
		fmt.Println("Index:", index, "Value:", value)
		mySlice = append(mySlice, value+10)
	}
	fmt.Println("Original slice:", mySlice)
	fmt.Println("Copied slice:", copySlice)
}
