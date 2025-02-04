package main

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

const (
	numEntries   = 100000
	targetGender = "Female" // The gender to search for
)

type Gender struct {
	ID     int
	Gender string
}

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
func searchSlice(genders []Gender, target string) bool {
	for _, g := range genders {
		if g.Gender == target {
			return true
		}
	}
	return false
}
func searchMap(genderMap map[int]string, target string) bool {
	for _, g := range genderMap {
		if g == target {
			return true
		}
	}
	return false
}
func BenchmarkSliceSearch(b *testing.B) {
	genders, _ := initData()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = searchSlice(genders, targetGender)
	}
}
func BenchmarkMapSearch(b *testing.B) {
	_, genderMap := initData()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = searchMap(genderMap, targetGender)
	}
}
func main() {
	fmt.Println("Run `go test -bench=.` to see the benchmark results")
}
