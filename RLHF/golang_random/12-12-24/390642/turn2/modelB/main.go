package main

import (
	"fmt"
	"math/rand"
	"time"
)

// SortStrategy defines the interface for different sorting strategies.
type SortStrategy interface {
	Sort([]int) []int
}

// BubbleSortStrategy implements the SortStrategy interface using the Bubble Sort algorithm.
type BubbleSortStrategy struct{}

func (b *BubbleSortStrategy) Sort(arr []int) []int {
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

// QuickSortStrategy implements the SortStrategy interface using the Quick Sort algorithm.
type QuickSortStrategy struct{}

func (q *QuickSortStrategy) Sort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	pivot := arr[len(arr)/2]
	left := []int{}
	middle := []int{}
	right := []int{}
	for _, v := range arr {
		if v < pivot {
			left = append(left, v)
		} else if v > pivot {
			right = append(right, v)
		} else {
			middle = append(middle, v)
		}
	}
	return append(q.Sort(left), middle...) + append([]int{}, q.Sort(right)...)
}

// Context holds a reference to a SortStrategy and provides a method to sort an array.
type Context struct {
	strategy SortStrategy
}

// SetStrategy sets the sorting strategy for the context.
func (c *Context) SetStrategy(strategy SortStrategy) {
	c.strategy = strategy
}

// Sort sorts an array using the current sorting strategy.
func (c *Context) Sort(arr []int) []int {
	return c.strategy.Sort(arr)
}

func main() {
	// Initialize random array
	rand.Seed(time.Now().UnixNano())
	arr := make([]int, 10)
	for i := range arr {
		arr[i] = rand.Intn(100)
	}
	fmt.Println("Original array:", arr)

	// Create a context
	context := &Context{}

	// Set the strategy to Bubble Sort and sort
	bubbleSort := &BubbleSortStrategy{}
	context.SetStrategy(bubbleSort)
	fmt.Println("\nSorted with Bubble Sort:", context.Sort(arr))

	// Set the strategy to Quick Sort and sort
	quickSort := &QuickSortStrategy{}
	context.SetStrategy(quickSort)
	fmt.Println("\nSorted with Quick Sort:", context.Sort(arr))
}
