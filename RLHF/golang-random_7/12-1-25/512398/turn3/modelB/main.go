package main

import (
	"fmt"
	"sort"
)

type MySlice[T sort.Interface] []T

// Sorting
func (s MySlice[T]) Len() int           { return len(s) }
func (s MySlice[T]) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s MySlice[T]) Less(i, j int) bool { return s[i].Less(s[j]) }

// Searching
func (s MySlice[T]) Search(target T) int {
	i, j := 0, len(s)-1
	for i <= j {
		h := int(uint(i+j) >> 1) // avoid overflow when computing h
		if s[h].Less(target) {
			i = h + 1
		} else if target.Less(s[h]) {
			j = h - 1
		} else {
			return h
		}
	}
	return -1
}

func main() {
	// Custom sorting for strings
	type MyStringSlice MySlice[string]
	sort.Sort(MyStringSlice{"apple", "banana", "cherry", "orange", "grape"})

	// Custom sorting for integers
	type MyIntSlice MySlice[int]
	sort.Sort(MyIntSlice{3, 1, 4, 1, 5, 9, 2, 6})

	//Using the custom slices
	fruitSlice := MyStringSlice{"apple", "banana", "cherry", "orange", "grape"}
	result := fruitSlice.Search("cherry")
	if result == -1 {
		fmt.Println("Element not found")
	} else {
		fmt.Println("Element found at index:", result)
	}

	intSlice := MyIntSlice{3, 1, 4, 1, 5, 9, 2, 6}
	target := 4
	result = intSlice.Search(target)
	if result == -1 {
		fmt.Println("Element", target, "not found.")
	} else {
		fmt.Println("Element", target, "found at index:", result)
	}
}
