package main

import "fmt"

func main() {
	mySlice := []int{1, 2, 3, 4, 5}

	for i := 0; i < len(mySlice); i++ {
		fmt.Println("Current Value:", mySlice[i])

		// Example modification: Remove even numbers
		if mySlice[i]%2 == 0 {
			// Remove the current element by appending the slice up to the current index
			// with the slice after the current index
			mySlice = append(mySlice[:i], mySlice[i+1:]...)

			// Decrement index to account for the removed element
			i--
		} else {
			// Example modification: Double odd numbers
			mySlice[i] *= 2
		}
	}

	fmt.Println("Modified slice:", mySlice)
}
