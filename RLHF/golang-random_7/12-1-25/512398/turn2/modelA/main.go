package main

import (
	"fmt"
	"sort"
	"strings"
)

// mySlice is a custom slice type with sorting and searching methods.
type mySlice []string

// Len is part of the sort.Interface.
func (s mySlice) Len() int {
	return len(s)
}

// Swap is part of the sort.Interface.
func (s mySlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// Less is part of the sort.Interface.
func (s mySlice) Less(i, j int) bool {
	// We can define any sorting logic here.
	// In this example, sort in reverse alphabetical order.
	return strings.Compare(s[i], s[j]) > 0
}

// BinarySearch performs a binary search on the sorted slice.
func (s mySlice) BinarySearch(target string) int {
	low := 0
	high := len(s) - 1

	for low <= high {
		mid := low + (high-low)/2
		// Adjust the comparison based on the sorting logic
		if strings.Compare(s[mid], target) > 0 {
			high = mid - 1
		} else if strings.Compare(s[mid], target) < 0 {
			low = mid + 1
		} else {
			return mid
		}
	}

	return -1
}

func main() {
	// Create a new instance of mySlice and add some strings
	fruits := mySlice{"banana", "apple", "cherry", "date", "grape"}

	// Sort the slice using the custom Less method
	sort.Sort(fruits)

	// Print the sorted slice
	fmt.Println("Sorted slice:", fruits)

	// Perform a binary search for a specific element
	target := "cherry"
	index := fruits.BinarySearch(target)

	if index != -1 {
		fmt.Println("Element found at index:", index)
	} else {
		fmt.Println("Element not found")
	}
}
