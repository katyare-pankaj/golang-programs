package main

import (
	"fmt"
	"sort"
)

type IntSlice []int

func (s IntSlice) Len() int           { return len(s) }
func (s IntSlice) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s IntSlice) Less(i, j int) bool { return s[i] < s[j] }

func (s IntSlice) BinarySearch(target int) int {
	i, j := 0, len(s)-1
	for i <= j {
		h := int(uint(i+j) >> 1) // avoid overflow when computing h
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
	var intSlice IntSlice
	intSlice = append(intSlice, 10, 20, 30, 40, 50)

	sort.Sort(intSlice)
	fmt.Println("Sorted slice:", intSlice)

	target := 30
	result := intSlice.BinarySearch(target)
	if result == -1 {
		fmt.Println("Element", target, "not found.")
	} else {
		fmt.Println("Element", target, "found at index:", result)
	}
}
