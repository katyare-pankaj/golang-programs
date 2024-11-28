package main

import (
	"fmt"
	"sync"
)

func mergeSortConcurrent(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	mid := len(arr) / 2
	left := arr[:mid]
	right := arr[mid:]

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		left = mergeSortConcurrent(left)
		wg.Done()
	}()

	go func() {
		right = mergeSortConcurrent(right)
		wg.Done()
	}()

	wg.Wait()

	return merge(left, right)
}

func merge(left, right []int) []int {
	var result []int
	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	result = append(result, left[i:]...)
	result = append(result, right[j:]...)

	return result
}

func main() {
	arr := []int{9, 14, 3, 2, 43, 11, 58, 22}
	fmt.Println("Original array:", arr)

	sortedArray := mergeSortConcurrent(arr)
	fmt.Println("Sorted array:", sortedArray)
}
