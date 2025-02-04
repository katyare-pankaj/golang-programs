package main

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

const numEntries = 100000

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

// Benchmark for slice
func BenchmarkSlice(b *testing.B) {
	genders, _ := initData()
	for i := 0; i < b.N; i++ {
		// Retrieve from slice
		for j := 0; j < numEntries; j++ {
			_ = genders[j].Gender
		}
	}
}

// Benchmark for map
func BenchmarkMap(b *testing.B) {
	_, genderMap := initData()
	for i := 0; i < b.N; i++ {
		// Retrieve from map
		for j := 0; j < numEntries; j++ {
			_ = genderMap[j]
		}
	}
}

// Benchmark adding new entries
func BenchmarkAddSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		genders := make([]Gender, 0, numEntries)
		for j := 0; j < numEntries; j++ {
			genders = append(genders, Gender{ID: j, Gender: "Male"})
		}
	}
}

func BenchmarkAddMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		genderMap := make(map[int]string, numEntries)
		for j := 0; j < numEntries; j++ {
			genderMap[j] = "Male"
		}
	}
}

// Running benchmarks
func main() {
	fmt.Println("Run `go test -bench=.` to see the benchmark results")
}
