package main

import (
	"fmt"
	"strings"
	"time"
)

const numIterations = 1000000

func main() {
	start := time.Now()
	concatenateUsingPlus()
	fmt.Println("Concatenation using +:", time.Since(start))
	start = time.Now()
	concatenateUsingBuilder()
	fmt.Println("Concatenation using strings.Builder:", time.Since(start))
	start = time.Now()
	concatenateUsingSprintf()
	fmt.Println("Concatenation using fmt.Sprintf:", time.Since(start))
}
func concatenateUsingPlus() {
	var s string
	for i := 0; i < numIterations; i++ {
		s = s + "Hello, " + fmt.Sprintf("%d", i) + "!"
	}
}
func concatenateUsingBuilder() {
	var b strings.Builder
	for i := 0; i < numIterations; i++ {
		b.WriteString("Hello, ")
		b.WriteString(fmt.Sprintf("%d", i))
		b.WriteString("!")
	}
}
func concatenateUsingSprintf() {
	for i := 0; i < numIterations; i++ {
		s := fmt.Sprintf("Hello, %d!", i)
	}
}
