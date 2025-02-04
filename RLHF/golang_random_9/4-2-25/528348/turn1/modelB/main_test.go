package main

import (
	"fmt"
	"math/rand"
	"testing"
)

const (
	numEntries = 1000000
	numGenders = 3
)

var (
	genders []string = []string{"Male", "Female", "Non-binary"}
)

func generateGenderData() []string {
	data := make([]string, numEntries)
	for i := 0; i < numEntries; i++ {
		data[i] = genders[rand.Intn(numGenders)]
	}
	return data
}

func BenchmarkMapGenderData(b *testing.B) {
	for n := 0; n < b.N; n++ {
		data := generateGenderData()
		genderCount := make(map[string]int)

		for _, gender := range data {
			genderCount[gender]++
		}
	}
}

func BenchmarkSliceGenderData(b *testing.B) {
	for n := 0; n < b.N; n++ {
		data := generateGenderData()
		genderCount := make([]int, numGenders)
		genderMap := map[string]int{
			"Male":       0,
			"Female":     1,
			"Non-binary": 2,
		}

		for _, gender := range data {
			genderCount[genderMap[gender]]++
		}
	}
}

func main() {
	// Generating the data
	data := generateGenderData()

	// Ease of use and Scalability comparison
	genderCount := make(map[string]int)
	for _, gender := range data {
		genderCount[gender]++
	}
	fmt.Println("Gender Count using Map:")
	for gender, count := range genderCount {
		fmt.Printf("%s: %d\n", gender, count)
	}

	genderCount = make([]int, numGenders)
	genderMap := map[string]int{
		"Male":       0,
		"Female":     1,
		"Non-binary": 2,
	}
	for _, gender := range data {
		genderCount[genderMap[gender]]++
	}
	fmt.Println("\nGender Count using Slice:")
	for i, count := range genderCount {
		fmt.Printf("%s: %d\n", genders[i], count)
	}
}
