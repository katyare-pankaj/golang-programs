package main

import (
	"fmt"
	"sort"
)

func rangePartitionWithOverlap(data []int, numPartitions int, overlap int) [][]int {
	sort.Ints(data)

	partitionSize := (len(data) + numPartitions - 1) / numPartitions
	startIndex := 0
	partitions := make([][]int, 0)

	for i := 0; i < numPartitions; i++ {
		endIndex := startIndex + partitionSize - 1
		if endIndex >= len(data) {
			endIndex = len(data) - 1
		}
		partitions = append(partitions, data[startIndex:endIndex+1])
		startIndex = endIndex - overlap + 1
	}
	return partitions
}

func main() {
	data := []int{1, 12, 5, 23, 7, 8, 3, 18, 2, 4, 6, 11, 20} // Large dataset
	numPartitions := 3
	overlap := 2 // Adjust the overlap as needed
	partitions := rangePartitionWithOverlap(data, numPartitions, overlap)

	for i, partition := range partitions {
		fmt.Printf("Partition %d: %v\n", i, partition)
	}
}
