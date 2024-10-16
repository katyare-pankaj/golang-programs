package main

import (
	"fmt"
)

// quicksort function
func quicksort(arr []int, low int, high int) {
	if low < high {
		p := partition(arr, low, high)
		quicksort(arr, low, p-1)
		quicksort(arr, p+1, high)
	}
}

// partition function
func partition(arr []int, low int, high int) int {
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

// Insertion sort function
func insertionSort(arr []int, low int, high int) {
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

// HybridSort function:  uses quicksort for large arrays and insertion sort for small arrays.
func hybridSort(arr []int, low int, high int, threshold int) {
	if low < high {
		size := high - low + 1
		if size <= threshold {
			insertionSort(arr, low, high)
		} else {
			p := partition(arr, low, high)
			hybridSort(arr, low, p-1, threshold)
			hybridSort(arr, p+1, high, threshold)
		}
	}
}

func main() {
	arr := []int{8, 3, 2, 5, 4, 7, 6, 1}
	fmt.Println("Before sorting: ", arr)
	hybridSort(arr, 0, len(arr)-1, 4) //threshold set to 4, you can change it
	fmt.Println("After sorting: ", arr)
}
