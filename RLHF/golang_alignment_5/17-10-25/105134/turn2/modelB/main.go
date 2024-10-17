package main

import (
	"fmt"
	"math/rand"
	"time"
)

const threshold = 10 // Threshold for switching to insertion sort

// hybridSort function: uses quicksort for large arrays and insertion sort for small arrays.
func hybridSort(arr []int, low int, high int) {
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

// partition function: This now handles the pivot and element swaps correctly.
func partition(arr []int, low int, high int) int {
	pivot := arr[high]
	i := low

	for j := low; j < high; j++ {
		if arr[j] < pivot { // Ensure we're comparing correctly
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[high] = arr[high], arr[i]
	return i
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

// Generate random array of sensor data
func generateSensorData(size int) []int {
	data := make([]int, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		data[i] = rand.Intn(1000) + 1 // Generate random integers between 1 and 1000 to avoid zeros
	}
	return data
}

func main() {
	sensorDataSize := 100000 // Change the size as needed
	sensorData := generateSensorData(sensorDataSize)

	fmt.Println("Before sorting (first 10 elements): ", sensorData[0:10]) // Print first 10 elements

	startTime := time.Now()
	hybridSort(sensorData, 0, len(sensorData)-1)
	endTime := time.Now()

	fmt.Println("After sorting (first 10 elements): ", sensorData[0:10]) // Print first 10 elements
	fmt.Println("Sorting time: ", endTime.Sub(startTime))
}
