package main

import (
	"fmt"
	"sort"
)

// DataPoint represents a point in the dataset with a float feature and a label.
type DataPoint struct {
	Feature float64
	Label   string
}

func main() {
	// Sample dataset
	data := []DataPoint{
		{Feature: 2.3, Label: "B"},
		{Feature: 1.1, Label: "A"},
		{Feature: 4.5, Label: "D"},
		{Feature: 3.0, Label: "C"},
	}

	// Sort by Feature
	sort.Slice(data, func(i, j int) bool {
		return data[i].Feature < data[j].Feature
	})
	fmt.Println("Sorted data:", data)

	// Perform binary search for a specific feature (after sorting)
	target := 3.0
	index := sort.Search(len(data), func(i int) bool {
		return data[i].Feature >= target
	})

	if index < len(data) && data[index].Feature == target {
		fmt.Printf("Found data point with feature %v at index %d: %+v\n", target, index, data[index])
	} else {
		fmt.Printf("No data point found with feature %v\n", target)
	}
}
