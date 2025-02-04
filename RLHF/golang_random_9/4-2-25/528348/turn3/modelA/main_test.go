package main

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

const (
	numEntries   = 100000
	updateGender = "Other" // The gender value to update to
)

type Gender struct {
	ID     int
	Gender string
}

// Initialize data for benchmarking
func initData() ([]Gender, map[int]string) {
	slice := make([]Gender, 0, numEntries)
	m := make(map[int]string, numEntries)

	genders := []string{"Male", "Female", "Non-Binary"}
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < numEntries; i++ {
		g := genders[rand.Intn(len(genders))]
		slice = append(slice, Gender{ID: i, Gender: g})
		m[i] = g
	}

	return slice, m
}

// Update gender values in a slice
func updateSlice(genders []Gender, newGender string) {
	for i := range genders {
		genders[i].Gender = newGender
	}
}

// Update gender values in a map
func updateMap(genderMap map[int]string, newGender string) {
	for i := range genderMap {
		genderMap[i] = newGender
	}
}

// Benchmark updating slice
func BenchmarkSliceUpdate(b *testing.B) {
	genders, _ := initData()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		updateSlice(genders, updateGender)
	}
}

// Benchmark updating map
func BenchmarkMapUpdate(b *testing.B) {
	_, genderMap := initData()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		updateMap(genderMap, updateGender)
	}
}

func main() {
	fmt.Println("Run `go test -bench=.` to see the benchmark results")
}
