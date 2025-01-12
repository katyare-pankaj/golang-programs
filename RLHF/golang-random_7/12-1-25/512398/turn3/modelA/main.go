package main

import (
	"fmt"
	"sort"
)

type CustomIntSlice []int

// Define sorting logic using sort.Interface
func (s CustomIntSlice) Len() int {
	return len(s)
}

func (s CustomIntSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s CustomIntSlice) Less(i, j int) bool {
	// Example: Sort in ascending order
	return s[i] < s[j]
}

// Define binary search logic
func (s CustomIntSlice) BinarySearch(target int) int {
	i, j := 0, len(s)-1
	for i <= j {
		h := int(uint(i+j) >> 1) // Avoid overflow
		if s[h] < target {
			i = h + 1
		} else if s[h] > target {
			j = h - 1
		} else {
			return h
		}
	}
	return -1
}

func main() {
	// Create an instance of the custom slice
	slice := CustomIntSlice{5, 10, 3, 8, 7}

	// Sort the slice using the defined sorting logic
	sort.Sort(slice)

	// Print the sorted slice
	fmt.Println("Sorted slice:", slice)

	// Search for a target value in the sorted slice
	target := 7
	index := slice.BinarySearch(target)

	// Print the result of the search
	if index == -1 {
		fmt.Println("Element", target, "not found.")
	} else {
		fmt.Println("Element", target, "found at index:", index)
	}
}
