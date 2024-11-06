package main

import (
	"fmt"
	"hash/fnv"
)

func hashPartition(data []int, numPartitions int) [][]int {
	partitions := make([][]int, numPartitions)

	for _, value := range data {
		h := fnv.New32a()
		_, _ = h.Write([]byte(fmt.Sprint(value)))
		partitionIndex := int(h.Sum32()) % numPartitions
		partitions[partitionIndex] = append(partitions[partitionIndex], value)
	}
	return partitions
}

func main() {
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12} // Large dataset
	numPartitions := 4
	partitions := hashPartition(data, numPartitions)

	for i, partition := range partitions {
		fmt.Printf("Partition %d: %v\n", i, partition)
	}
}
