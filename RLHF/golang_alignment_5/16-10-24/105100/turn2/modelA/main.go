package main

import (
	"fmt"
	"math/rand"
	"time"
)

const threshold = 10 // Threshold for switching to insertion sort

func hybridSort(arr []int, low, high int) {
	if low < high {
		size := high - low + 1
		if size <= threshold {
			insertionSort(arr, low, high)
		} else {
			p := partition(arr, low, high)
			hybridSort(arr, low, p-1)
			hybridSort(arr, p+1, high)
		}
	}
}

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

func generateRandomArray(size int) []int {
	arr := make([]int, size)
	for i := 0; i < size; i++ {
		arr[i] = rand.Intn(1000)
	}
	return arr
}

func main() {
	rand.Seed(time.Now().UnixNano())

	arrSize := 10000
	arr := generateRandomArray(arrSize)

	fmt.Println("Before sorting:")
	fmt.Println(arr)

	startTime := time.Now()
	hybridSort(arr, 0, len(arr)-1)
	endTime := time.Now()

	fmt.Println("\nAfter sorting:")
	fmt.Println(arr)

	fmt.Println("\nExecution time:", endTime.Sub(startTime))
}
