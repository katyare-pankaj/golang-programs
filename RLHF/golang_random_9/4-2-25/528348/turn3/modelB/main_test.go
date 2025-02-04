package main

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

const (
	numEntries   = 100000
	updateGender = "Non-Binary" // The new gender to update
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

func updateSlice(genders []Gender, updateGender string) {
	for i, g := range genders {
		if g.Gender == "Female" {
			genders[i].Gender = updateGender
		}
	}
}

func updateMap(genderMap map[int]string, updateGender string) {
	for i, g := range genderMap {
		if g == "Female" {
			genderMap[i] = updateGender
		}
	}
}

func BenchmarkSliceUpdate(b *testing.B) {
	genders, _ := initData()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		updateSlice(genders, updateGender)
	}
}

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
