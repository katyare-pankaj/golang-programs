package main

import (
	"fmt"
	"math/rand"
	"testing"
)

func reverseSliceForLoop(s []int) {
	for i := 0; i < len(s)/2; i++ {
		s[i], s[len(s)-i-1] = s[len(s)-i-1], s[i]
	}
}

func reverseSliceCopy(s []int) {
	temp := make([]int, len(s))
	copy(temp, s)
	for i := 0; i < len(s); i++ {
		s[i] = temp[len(s)-i-1]
	}
}

func BenchmarkReverseSliceForLoop(b *testing.B) {
	sliceSize := 100000
	slice := randSlice(sliceSize)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		reverseSliceForLoop(slice)
	}
}

func BenchmarkReverseSliceCopy(b *testing.B) {
	sliceSize := 100000
	slice := randSlice(sliceSize)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		reverseSliceCopy(slice)
	}
}
func randSlice(size int) []int {
	slice := make([]int, size)
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(1000000)
	}
	return slice
}

func main() {
	fmt.Println("Before reverse:")
	fmt.Println(randSlice(10))

	fmt.Println("\nReversing slice using for loop...")
	sl1 := randSlice(100000)
	reverseSliceForLoop(sl1)
	fmt.Println(sl1[0:10]) //Printing only first 10 elements for brevity

	fmt.Println("\nReversing slice using copy function...")
	sl2 := randSlice(100000)
	reverseSliceCopy(sl2)
	fmt.Println(sl2[0:10]) //Printing only first 10 elements for brevity

}
