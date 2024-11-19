package main

import "fmt"

// Define the behavior interface for sorting algorithms
type SortingAlgorithm interface {
	Sort([]int) []int
}

// BubbleSort implements the SortingAlgorithm interface
type BubbleSort struct{}

func (b BubbleSort) Sort(arr []int) []int {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

// MergeSort implements the SortingAlgorithm interface
type MergeSort struct{}

func (m MergeSort) Sort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	mid := len(arr) / 2
	left := m.Sort(arr[:mid])
	right := m.Sort(arr[mid:])
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
	for i < len(left) {
		result = append(result, left[i])
		i++
	}
	for j < len(right) {
		result = append(result, right[j])
		j++
	}
	return result
}

// Module that uses sorting algorithms via the interface
type SortingModule struct {
	Algorithm SortingAlgorithm
}

func (s *SortingModule) SortArray(arr []int) []int {
	return s.Algorithm.Sort(arr)
}

func main() {
	// Create an instance of the SortingModule with BubbleSort
	bubbleSortModule := SortingModule{Algorithm: BubbleSort{}}
	unsortedArray1 := []int{64, 34, 25, 12, 22, 11, 90}
	sortedArray1 := bubbleSortModule.SortArray(unsortedArray1)
	fmt.Println("BubbleSort: ", sortedArray1) // Output: BubbleSort: [11 12 22 25 34 64 90]

	// Create an instance of the SortingModule with MergeSort
	mergeSortModule := SortingModule{Algorithm: MergeSort{}}
	unsortedArray2 := []int{64, 34, 25, 12, 22, 11, 90}
	sortedArray2 := mergeSortModule.SortArray(unsortedArray2)
	fmt.Println("MergeSort: ", sortedArray2) // Output: MergeSort: [11 12 22 25 34 64 90]
}
