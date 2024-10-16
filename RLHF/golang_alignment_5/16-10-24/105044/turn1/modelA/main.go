package main

import "fmt"

const threshold = 10 // Choose an appropriate threshold value

// Quicksort implementation
func quicksort(arr []int, low, high int) {
	if low < high {
		pi := partition(arr, low, high)
		quicksort(arr, low, pi-1)
		quicksort(arr, pi+1, high)
	}
}

func partition(arr []int, low, high int) int {
	pivot := arr[high]
	i := low - 1
	for j := low; j < high; j++ {
		if arr[j] <= pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}

// Insertion sort implementation
func insertionSort(arr []int, low, high int) {
	for i := low + 1; i <= high; i++ {
		key := arr[i]
		j := i - 1
		for j >= low && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
}

// Hybrid quicksort with insertion sort for smaller sections
func hybridQuicksort(arr []int, low, high int) {
	if high-low+1 <= threshold {
		insertionSort(arr, low, high)
	} else {
		quicksort(arr, low, high)
	}
}

func main() {
	arr := []int{8, 3, 2, 5, 4, 7, 6, 1}
	fmt.Println("Unsorted array:", arr)

	hybridQuicksort(arr, 0, len(arr)-1)

	fmt.Println("Sorted array: ", arr)
}
