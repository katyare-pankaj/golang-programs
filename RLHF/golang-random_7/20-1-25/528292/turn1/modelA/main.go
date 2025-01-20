package main

import (
	"fmt"
	"sort"
)

// DataPoint represents a single data point in machine learning.
type DataPoint struct {
	Feature float64
	Label   string
}

// ByFeature implements sort.Interface based on the Feature field.
type ByFeature []DataPoint

func (a ByFeature) Len() int           { return len(a) }
func (a ByFeature) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByFeature) Less(i, j int) bool { return a[i].Feature < a[j].Feature }

// SortData sorts the data points by the Feature field.
func SortData(data []DataPoint) {
	sort.Sort(ByFeature(data))
}

// SearchData performs a binary search for a specific feature value.
// It assumes the slice is already sorted by Feature.
// It returns the index of the data point with the matching feature, or -1 if not found.
func SearchData(data []DataPoint, target float64) int {
	i := sort.Search(len(data), func(i int) bool {
		return data[i].Feature >= target
	})
	if i < len(data) && data[i].Feature == target {
		return i
	}
	return -1
}

func main() {
	// Sample data points
	data := []DataPoint{
		{Feature: 2.3, Label: "B"},
		{Feature: 1.1, Label: "A"},
		{Feature: 4.5, Label: "D"},
		{Feature: 3.0, Label: "C"},
	}

	// Sort the data by Feature
	SortData(data)
	fmt.Println("Sorted data:", data)

	// Search for a specific feature
	target := 3.0
	index := SearchData(data, target)

	if index != -1 {
		fmt.Printf("Found data point with feature %v at index %d: %+v\n", target, index, data[index])
	} else {
		fmt.Printf("No data point found with feature %v\n", target)
	}
}
