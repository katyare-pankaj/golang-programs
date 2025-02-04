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

// Benchmark for searching in slice
func BenchmarkSearchSlice(b *testing.B) {
	genders, _ := initData()
	targetGender := "Female"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, g := range genders {
			if g.Gender == targetGender {
				break
			}
		}
	}
}

// Benchmark for searching in map
func BenchmarkSearchMap(b *testing.B) {
	_, genderMap := initData()
	targetGender := "Female"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, g := range genderMap {
			if g == targetGender {
				break
			}
		}
	}
}

// Running benchmarks
func main() {
	fmt.Println("Run `go test -bench=.` to see the benchmark results")
}
