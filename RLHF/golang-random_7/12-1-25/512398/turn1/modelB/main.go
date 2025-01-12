package main

import (
	"fmt"
	"sort"
	"strings"
)

type mySlice []string

// sorting
func (s mySlice) Len() int      { return len(s) }
func (s mySlice) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s mySlice) Less(i, j int) bool {
	return strings.Compare(s[i], s[j]) < 0
}

// searching
func (s mySlice) Search(x string) int {
	i, j := 0, len(s)-1
	for i <= j {
		h := int(uint(i+j) >> 1) // avoid overflow when computing h
		// Adjust comparison based on your sorting logic
		if strings.Compare(s[h], x) < 0 {
			i = h + 1
		} else if strings.Compare(s[h], x) > 0 {
			j = h - 1
		} else {
			return h
		}
	}
	return -1
}

func main() {
	var fruitSlice mySlice
	fruitSlice = append(fruitSlice, "apple", "banana", "cherry", "orange", "grape")
	sort.Sort(fruitSlice) // Sort the slice using the custom sorting algorithm
	fmt.Println("Sorted slice:", fruitSlice)
	result := fruitSlice.Search("cherry")
	if result == -1 {
		fmt.Println("Element not found")
	} else {
		fmt.Println("Element found at index:", result)
	}
}
