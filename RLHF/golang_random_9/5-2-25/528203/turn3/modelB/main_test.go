package main

import (
	"testing"
	"time"
)

func inefficientFunction(data []int) int {
	sum := 0
	for _, value := range data {
		// Simulated heavy calculation
		time.Sleep(10 * time.Millisecond)
		sum += value
	}
	return sum
}
func efficientFunction(data []int) int {
	sum := 0
	for _, value := range data {
		sum += value
	}
	return sum
}
func BenchmarkInefficientFunction(b *testing.B) {
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i := 0; i < b.N; i++ {
		b.ReportAllocs()
		inefficientFunction(data)
	}
}
func BenchmarkEfficientFunction(b *testing.B) {
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i := 0; i < b.N; i++ {
		b.ReportAllocs()
		efficientFunction(data)
	}
}
