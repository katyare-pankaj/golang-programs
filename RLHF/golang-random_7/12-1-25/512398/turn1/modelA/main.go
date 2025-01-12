package main

import (
	"fmt"
)

// CustomSlice represents a custom slice type with sorting and searching methods.
type CustomSlice []int

// NewCustomSlice creates a new instance of CustomSlice.
func NewCustomSlice(elements ...int) CustomSlice {
	return CustomSlice(elements)
}

// Sort sorts the CustomSlice in place using a custom sorting function.
func (cs *CustomSlice) Sort(sortFunc func(int, int) bool) {
	for i := range *cs {
		for j := i + 1; j < len(*cs); j++ {
			if !sortFunc((*cs)[i], (*cs)[j]) {
				(*cs)[i], (*cs)[j] = (*cs)[j], (*cs)[i]
			}
		}
	}
}

// Search performs a linear search on the CustomSlice.
func (cs CustomSlice) Search(target int) int {
	for i, v := range cs {
		if v == target {
			return i
		}
	}
	return -1
}

// BinarySearch performs a binary search on the CustomSlice, assuming it's sorted.
func (cs CustomSlice) BinarySearch(target int) int {
	left, right := 0, len(cs)-1
	for left <= right {
		mid := left + (right-left)/2
		if cs[mid] == target {
			return mid
		} else if cs[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

func main() {
	// Create a CustomSlice with some elements
	slice := NewCustomSlice(3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5)

	fmt.Println("Original slice:", slice)

	// Sort the slice using a custom sort function
	slice.Sort(func(a, b int) bool {
		return a > b
	})

	fmt.Println("Sorted slice (descending):", slice)

	// Search for an element
	index := slice.Search(9)
	if index != -1 {
		fmt.Printf("Found 9 at index: %d\n", index)
	} else {
		fmt.Println("9 not found")
	}

	// Perform a binary search
	index = slice.BinarySearch(5)
	if index != -1 {
		fmt.Printf("Found 5 at index: %d\n", index)
	} else {
		fmt.Println("5 not found")
	}
}
