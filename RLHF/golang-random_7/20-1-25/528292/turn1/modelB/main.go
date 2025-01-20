package main

import (
	"fmt"
	"sort"
)

// MLData represents a slice of float64
type MLData []float64

// AscendingOrder is a custom sorting order for MLData
type AscendingOrder MLData

func (a AscendingOrder) Len() int           { return len(a) }
func (a AscendingOrder) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a AscendingOrder) Less(i, j int) bool { return a[i] < a[j] }

func main() {
	// Sample ML data
	data := MLData{12.3, 5.4, 3.2, 7.8, 10.1, 9.5, 4.6}

	//Sort data using the custom AscendingOrder sorting function
	sort.Sort(AscendingOrder(data))
	fmt.Println("Sorted data: ", data)

	//Binary search for an element
	elementToFind := 7.8
	index := binarySearch(data, elementToFind)
	if index != -1 {
		fmt.Println("Element found at index: ", index)
	} else {
		fmt.Println("Element not found.")
	}

	//Quickselect for an element by rank
	elementToSelect := 2 //2nd smallest element
	result := quickselect(data, elementToSelect)
	fmt.Println("\nElement at rank", elementToSelect, "is: ", result)

}

// binarySearch performs a binary search on the sorted MLData slice.
func binarySearch(data MLData, target float64) int {
	low := 0
	high := len(data) - 1

	for low <= high {
		mid := (low + high) / 2
		if data[mid] == target {
			return mid
		} else if data[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

// quickselect performs the quickselect algorithm on the MLData slice to find the element at the 'rank' position.
func quickselect(data MLData, rank int) float64 {
	//Base case: If rank is negative or greater than the data length, return -1
	if rank < 1 || rank > len(data) {
		return -1.0
	}
	return quickselectRec(data, rank-1)
}

func quickselectRec(data MLData, rank int) float64 {
	//Pivot selection as the median of the first, middle, and last elements
	pivotIndex := medianIndex(data)
	pivot := data[pivotIndex]

	//Partition the data around the pivot
	left, right := partition(data, pivot)

	//Calculate the size of the left partition
	leftSize := len(left)

	//If the pivot is at the desired rank, return it
	if leftSize == rank {
		return pivot
	}

	//Recur for the left or right partition based on the rank
	if leftSize > rank {
		return quickselectRec(left, rank)
	} else {
		return quickselectRec(right, rank-leftSize-1)
	}
}

func medianIndex(data MLData) int {
	length := len(data)
	if length == 0 {
		return -1
	}
	return length / 2
}

func partition(data MLData, pivot float64) ([]float64, []float64) {
	var left, right []float64
	for _, value := range data {
		if value < pivot {
			left = append(left, value)
		} else {
			right = append(right, value)
		}
	}
	return left, right
}
