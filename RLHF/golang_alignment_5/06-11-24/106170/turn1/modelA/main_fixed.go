package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

// DataPartitioner interface defines the methods for data partitioning
type DataPartitioner interface {
	Partition(data []int, numPartitions int) [][]int
}

// HashPartitioner partitions data based on a hash function
type HashPartitioner struct{}

func (h HashPartitioner) Partition(data []int, numPartitions int) [][]int {
	partitions := make([][]int, numPartitions)

	for _, value := range data {
		partitionIndex := value % numPartitions
		partitions[partitionIndex] = append(partitions[partitionIndex], value)
	}
	return partitions
}

// RangePartitioner partitions data based on a predefined range
type RangePartitioner struct {
	minValue int
	maxValue int
}

func (r RangePartitioner) Partition(data []int, numPartitions int) [][]int {
	partitionSize := (r.maxValue-r.minValue)/numPartitions + 1
	partitions := make([][]int, numPartitions)

	for _, value := range data {
		partitionIndex := (value - r.minValue) / partitionSize
		partitions[partitionIndex] = append(partitions[partitionIndex], value)
	}
	return partitions
}

// SortedPartitioner partitions data by sorting it and then dividing it into equal parts
type SortedPartitioner struct{}

func (s SortedPartitioner) Partition(data []int, numPartitions int) [][]int {
	sort.Ints(data)
	partitionSize := len(data) / numPartitions
	partitions := make([][]int, numPartitions)

	for i := 0; i < len(data); i++ {
		partitionIndex := i / partitionSize
		if partitionIndex >= numPartitions {
			partitionIndex = numPartitions - 1
		}
		partitions[partitionIndex] = append(partitions[partitionIndex], data[i])
	}
	return partitions
}

func main() {
	// Generate sample data
	data := generateRandomData(100, 1, 10)
	numPartitions := 4

	// Test HashPartitioner
	hashPartitioner := HashPartitioner{}
	partitions := hashPartitioner.Partition(data, numPartitions)
	fmt.Println("Hash Partitioning:")
	printPartitions(partitions)

	// Test RangePartitioner
	minValue := min(data)
	maxValue := max(data)
	rangePartitioner := RangePartitioner{minValue, maxValue}
	partitions = rangePartitioner.Partition(data, numPartitions)
	fmt.Println("\nRange Partitioning:")
	printPartitions(partitions)

	// Test SortedPartitioner
	sortedPartitioner := SortedPartitioner{}
	partitions = sortedPartitioner.Partition(data, numPartitions)
	fmt.Println("\nSorted Partitioning:")
	printPartitions(partitions)
}

func generateRandomData(size int, minValue int, maxValue int) []int {
	data := make([]int, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		data[i] = rand.Intn(maxValue-minValue+1) + minValue
	}
	return data
}

func printPartitions(partitions [][]int) {
	for i, partition := range partitions {
		fmt.Printf("Partition %d: %v\n", i, partition)
	}
}

func min(data []int) int {
	if len(data) == 0 {
		return 0
	}
	minValue := data[0]
	for _, value := range data {
		if value < minValue {
			minValue = value
		}
	}
	return minValue
}

func max(data []int) int {
	if len(data) == 0 {
		return 0
	}
	maxValue := data[0]
	for _, value := range data {
		if value > maxValue {
			maxValue = value
		}
	}
	return maxValue
}
