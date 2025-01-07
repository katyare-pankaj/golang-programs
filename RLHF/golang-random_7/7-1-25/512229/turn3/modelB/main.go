package main

import (
	"fmt"
	"sort"
)

// Sortable interface defines the sorting methods
type Sortable interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}

// Sort function takes any Sortable slice and sorts it
func Sort(data Sortable) {
	sort.Sort(data)
}

// Define custom types for integers, strings, and a struct
type IntSlice []int
type StringSlice []string
type Person struct {
	Name string
	Age  int
}
type PersonSlice []Person

// Implement Sortable interface for each custom type
func (p IntSlice) Len() int           { return len(p) }
func (p IntSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p IntSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func (p StringSlice) Len() int           { return len(p) }
func (p StringSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p StringSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func (p PersonSlice) Len() int           { return len(p) }
func (p PersonSlice) Less(i, j int) bool { return p[i].Name < p[j].Name }
func (p PersonSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func main() {
	// Example usage:
	unsortedInts := []int{64, 34, 25, 12, 22, 11, 90}
	unsortedStrings := []string{"banana", "apple", "orange", "grape", "cherry"}
	unsortedPeople := []Person{{"Alice", 25}, {"Bob", 30}, {"Charlie", 22}}

	fmt.Println("Unsorted integers: ", unsortedInts)
	Sort(IntSlice(unsortedInts))
	fmt.Println("Sorted integers: ", unsortedInts)

	fmt.Println("Unsorted strings: ", unsortedStrings)
	Sort(StringSlice(unsortedStrings))
	fmt.Println("Sorted strings: ", unsortedStrings)

	fmt.Println("Unsorted people: ", unsortedPeople)
	Sort(PersonSlice(unsortedPeople))
	fmt.Println("Sorted people: ", unsortedPeople)
}
