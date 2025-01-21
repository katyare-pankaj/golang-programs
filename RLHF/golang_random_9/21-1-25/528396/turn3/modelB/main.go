package main

import (
	"fmt"
	"sort"
)

func mode(data []int) int {
	// Create a copy of the original slice to preserve it
	copyData := append([]int(nil), data...)
	sort.Ints(copyData)

	if len(copyData) == 0 {
		return 0
	}

	count := 1
	maxCount := 1
	mode := copyData[0]

	for i := 1; i < len(copyData); i++ {
		if copyData[i] == copyData[i-1] {
			count++
		} else {
			if count > maxCount {
				maxCount = count
				mode = copyData[i-1]
			}
			count = 1
		}
	}

	// Check for the last run of elements
	if count > maxCount {
		mode = copyData[len(copyData)-1]
	}

	return mode
}

func main() {
	// Original data slice (immutable)
	originalData := []int{1, 3, 5, 3, 7, 3, 5, 8, 1}

	fmt.Println("Original data:", originalData)

	// Calculate the mode of the data
	modeValue := mode(originalData)

	fmt.Println("Mode:", modeValue)

	// Confirm that the original data remains unchanged
	fmt.Println("Original data after mode calculation:", originalData)
}
