package main

import "fmt"

func main() {
	mySlice := []int{1, 2, 3, 4, 5}

	defer func() {
		// Perform all slice modifications after iteration
		for i, value := range modifiedIndices {
			mySlice = append(mySlice[:i], mySlice[i+1:]...)
		}
	}()

	modifiedIndices := []int{}

	for index, value := range mySlice {
		fmt.Println("Index:", index, "Value:", value)

		if value%2 == 0 {
			// Remember the indices of elements to be removed
			modifiedIndices = append(modifiedIndices, index)
		}

		// You can also modify elements safely here
		if value%3 == 0 {
			mySlice[index] *= 3
		}
	}

	fmt.Println("Original slice after iteration:", mySlice)
}
