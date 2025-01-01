package main

import (
	"fmt"
	"testing"
)

func fmtSprintfExample(a, b, c string) string {
	return fmt.Sprintf("%s%s%s", a, b, c)
}

func BenchmarkStringConcatenation(t *testing.B) {
	a := "Hello, "
	b := "world!"
	c := " How are you?"
	for i := 0; i < len(b); i++ {
		stringConcatenation(a, b, c)
	}
}

func BenchmarkFmtSprintfExample(t *testing.B) {
	a := "Hello, "
	b := "world!"
	c := " How are you?"
	for i := 0; i < len(b); i++ {
		fmtSprintfExample(a, b, c)
	}
}
