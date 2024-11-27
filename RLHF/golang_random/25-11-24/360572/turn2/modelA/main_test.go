// main.go
package main

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

type Person struct {
	Name string
	Age  int
	City string
}

func TestComparePersons(t *testing.T) {
	p1 := Person{
		Name: "Alice",
		Age:  30,
		City: "New York",
	}

	p2 := Person{
		Name: "Bob",
		Age:  25,
		City: "Los Angeles",
	}

	// Comparison using go-cmp
	if !cmp.Equal(p1, p2) {
		t.Fatalf("Persons not equal: %s", cmp.Diff(p1, p2))
	}

	fmt.Println("Comparison passed!")
}
