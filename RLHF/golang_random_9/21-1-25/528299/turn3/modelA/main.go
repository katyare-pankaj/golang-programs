package main

import "fmt"

func main() {
	mySlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	fmt.Println("Original slice:", mySlice)

	// Iterate using a traditional for loop with an index
	for i := 0; i < len(mySlice); {
		if mySlice[i]%2 == 0 {
			// Remove even numbers from the slice

			// Remove the current element by slicing
			mySlice = append(mySlice[:i], mySlice[i+1:]...)

			// Don't increment `i` here, because the elements have shifted to the left
		} else {
			// Only increment `i` if no element is removed
			// This ensures the next element is checked in the next iteration
			i++
		}
	}

	fmt.Println("Modified slice:", mySlice)
}
