package main

import (
	"bytes"
	"fmt"
	"testing"
)

func BenchmarkConcatenation(b *testing.B) {
	name := "Alice"
	age := 30
	for i := 0; i < b.N; i++ {
		_ = "Hello, my name is " + name + " and I am " + fmt.Sprintf("%d", age) + " years old."
	}
}

func BenchmarkSprintf(b *testing.B) {
	name := "Alice"
	age := 30
	for i := 0; i < b.N; i++ {
		_ = fmt.Sprintf("Hello, my name is %s and I am %d years old.", name, age)
	}
}

func BenchmarkBuffer(b *testing.B) {
	name := "Alice"
	age := 30
	buffer := &bytes.Buffer{}
	for i := 0; i < b.N; i++ {
		buffer.Reset()
		buffer.WriteString(fmt.Sprintf("Hello, my name is %s and I am %d years old.\n", name, age))
	}
}
