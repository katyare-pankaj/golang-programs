package main

import (
	"fmt"
	"testing"
)

func stringInterpolation(name string, age int) string {
	return fmt.Sprintf("Hello, my name is %s and I am %d years old.", name, age)
}

func fmtSprintfExample(name string, age int) string {
	return "Hello, my name is " + name + " and I am " + fmt.Sprintf("%d", age) + " years old."
}

func BenchmarkStringInterpolation(b *testing.B) {
	for i := 0; i < b.N; i++ {
		stringInterpolation("Alice", 25)
	}
}

func BenchmarkFmtSprintfExample(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fmtSprintfExample("Bob", 30)
	}
}

func main() {
	name := "Charlie"
	age := 27
	fmt.Println(stringInterpolation(name, age))
	fmt.Println(fmtSprintfExample(name, age))
}
