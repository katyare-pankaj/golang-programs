package main

import (
	"fmt"
)

// Sorter defines the interface for sorting strategies.
type Sorter interface {
	Sort(data []int)
}

// BubbleSort is a concrete strategy that implements bubble sort.
type BubbleSort struct{}

func (b *BubbleSort) Sort(data []int) {
	n := len(data)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if data[j] > data[j+1] {
				data[j], data[j+1] = data[j+1], data[j] // Swap
			}
		}
	}
}

// QuickSort is a concrete strategy that implements quick sort.
type QuickSort struct{}

func (q *QuickSort) Sort(data []int) {
	if len(data) < 2 {
		return
	}

	pivot := data[len(data)/2]
	left := []int{}
	right := []int{}

	for _, x := range data {
		if x < pivot {
			left = append(left, x)
		} else if x > pivot {
			right = append(right, x)
		}
	}
	q.Sort(left)
	q.Sort(right)
	copy(data, append(append(left, pivot), right...))
}

// MergeSort is a concrete strategy that implements merge sort.
type MergeSort struct{}

func (m *MergeSort) Sort(data []int) {
	if len(data) < 2 {
		return
	}
	mid := len(data) / 2
	left := make([]int, mid)
	right := make([]int, len(data)-mid)

	copy(left, data[:mid])
	copy(right, data[mid:])

	m.Sort(left)
	m.Sort(right)

	merge(data, left, right)
}

func merge(data, left, right []int) {
	i, j, k := 0, 0, 0
	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			data[k] = left[i]
			i++
		} else {
			data[k] = right[j]
			j++
		}
		k++
	}
	for i < len(left) {
		data[k] = left[i]
		i++
		k++
	}
	for j < len(right) {
		data[k] = right[j]
		j++
		k++
	}
}

// Context to use a sorting strategy.
type SortContext struct {
	sorter Sorter
}

// SetSorter sets the sorting strategy for the context.
func (c *SortContext) SetSorter(sorter Sorter) {
	c.sorter = sorter
}

// Sort executes the currently set sorting strategy.
func (c *SortContext) Sort(data []int) {
	c.sorter.Sort(data)
}

func main() {
	data := []int{5, 3, 4, 1, 2}

	// Create a context
	context := &SortContext{}

	// Set the strategy to BubbleSort and execute
	bubbleSort := &BubbleSort{}
	context.SetSorter(bubbleSort)
	fmt.Println("Using Bubble Sort:")
	context.Sort(data)
	fmt.Println(data)

	// Reset data for the next sort
	data = []int{5, 3, 4, 1, 2}

	// Set the strategy to QuickSort and execute
	quickSort := &QuickSort{}
	context.SetSorter(quickSort)
	fmt.Println("Using Quick Sort:")
	context.Sort(data)
	fmt.Println(data)

	// Reset data for the next sort
	data = []int{5, 3, 4, 1, 2}

	// Set the strategy to MergeSort and execute
	mergeSort := &MergeSort{}
	context.SetSorter(mergeSort)
	fmt.Println("Using Merge Sort:")
	context.Sort(data)
	fmt.Println(data)
}
