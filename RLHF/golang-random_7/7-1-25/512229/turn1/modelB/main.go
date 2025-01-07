package main

import (
	"fmt"
	"sort"
)

// Define an interface for sorting slices of any type
type Sortable interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}

// Implement the Sortable interface for a slice of integers
type IntSlice []int

func (p IntSlice) Len() int           { return len(p) }
func (p IntSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p IntSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

// Implement the Sortable interface for a slice of strings
type StringSlice []string

func (p StringSlice) Len() int           { return len(p) }
func (p StringSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p StringSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

// Implement the Sortable interface for a slice of structs with multiple fields
type Person struct {
	Name string
	Age  int
}
type PersonSlice []Person

func (p PersonSlice) Len() int { return len(p) }
func (p PersonSlice) Less(i, j int) bool {
	// Sort first by name, then by age
	if p[i].Name != p[j].Name {
		return p[i].Name < p[j].Name
	}
	return p[i].Age < p[j].Age
}
func (p PersonSlice) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

// Function to sort any slice that implements the Sortable interface
func Sort(data Sortable) {
	sort.Sort(data)
}

func main() {
	// Example 1: Sorting integers
	unsortedInts := []int{64, 34, 25, 12, 22, 11, 90}
	fmt.Println("Unsorted integers: ", unsortedInts)
	Sort(IntSlice(unsortedInts))
	fmt.Println("Sorted integers: ", unsortedInts)

	// Example 2: Sorting strings
	unsortedStrings := []string{"banana", "apple", "orange", "grape", "cherry"}
	fmt.Println("Unsorted strings: ", unsortedStrings)
	Sort(StringSlice(unsortedStrings))
	fmt.Println("Sorted strings: ", unsortedStrings)

	// Example 3: Sorting structs by multiple fields
	unsortedPeople := []Person{
		{"Alice", 25},
		{"Bob", 30},
		{"Charlie", 22},
		{"Alice", 28},
	}
	fmt.Println("Unsorted people: ", unsortedPeople)
	Sort(PersonSlice(unsortedPeople))
	fmt.Println("Sorted people: ", unsortedPeople)
}
