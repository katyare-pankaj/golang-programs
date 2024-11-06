package main

import (
	"fmt"
)

func main() {
	// Sample data
	data := []int{10, 20, 30, 40, 50, 60, 70, 80, 90}

	// Number of partitions
	numPartitions := 3

	// Partition the data based on range
	partitions := make([][]int, numPartitions)
	for _, value := range data {
		partitionIndex := (value / 10)
		if partitionIndex >= numPartitions {
			partitionIndex = numPartitions - 1
		}
		partitions[partitionIndex] = append(partitions[partitionIndex], value)
	}

	// Display the partitions
	for i, partition := range partitions {
		fmt.Printf("Partition %d: %v\n", i, partition)
	}
}
