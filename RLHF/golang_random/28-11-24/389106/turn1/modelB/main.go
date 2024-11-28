package main

import (
	"fmt"
	"sync"
)

func mergeSort(arr []int, wg *sync.WaitGroup) {
	if len(arr) <= 1 {
		wg.Done()
		return
	}

	mid := len(arr) / 2
	left := arr[:mid]
	right := arr[mid:]

	wg.Add(2)
	go mergeSort(left, wg)
	go mergeSort(right, wg)

	wg.Wait()
	merge(arr, left, right)
}

func merge(arr, left, right []int) {
	leftIndex := 0
	rightIndex := 0
	arrIndex := 0

	for leftIndex < len(left) && rightIndex < len(right) {
		if left[leftIndex] < right[rightIndex] {
			arr[arrIndex] = left[leftIndex]
			leftIndex++
		} else {
			arr[arrIndex] = right[rightIndex]
			rightIndex++
		}
		arrIndex++
	}

	for leftIndex < len(left) {
		arr[arrIndex] = left[leftIndex]
		leftIndex++
		arrIndex++
	}

	for rightIndex < len(right) {
		arr[arrIndex] = right[rightIndex]
		rightIndex++
		arrIndex++
	}
}

func main() {
	unsortedArray := []int{9, 14, 3, 2, 43, 11, 58, 22}
	fmt.Println("Unsorted array:", unsortedArray)

	wg := &sync.WaitGroup{}
	wg.Add(1)
	mergeSort(unsortedArray, wg)
	wg.Wait()

	fmt.Println("Sorted array: ", unsortedArray)
}
