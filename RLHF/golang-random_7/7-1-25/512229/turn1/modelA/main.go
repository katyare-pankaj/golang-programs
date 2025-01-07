package main

import (
	"fmt"
	"reflect"
	"sort"
)

// Sorter is an interface for elements that can be sorted.
type Sorter interface {
	Less(other Sorter) bool
	Equals(other Sorter) bool
}

// Implement Sorter for a simple struct
type Person struct {
	Name  string
	Age   int
	Email string
}

func (p Person) Less(other Sorter) bool {
	if otherPerson, ok := other.(Person); ok {
		return p.Name < otherPerson.Name
	}
	panic("not a Person")
}

func (p Person) Equals(other Sorter) bool {
	if otherPerson, ok := other.(Person); ok {
		return p.Name == otherPerson.Name
	}
	panic("not a Person")
}

// SortByField sorts a slice of Sorter elements by a specified field name.
func SortByField(slice sort.Slice, field string) {
	var rt reflect.TypeOf(slice[0])
	var fd reflect.Value

	for i := 0; i < rt.NumField(); i++ {
		if rt.Field(i).Name == field {
			fd = reflect.New(rt.Field(i).Type).Elem()
			break
		}
	}

	if fd.IsValid() {
		sort.SliceStable(slice, func(i, j int) bool {
			if slice[i].Equals(slice[j]) {
				return false // Do not change order if equal
			}

			iv := reflect.ValueOf(slice[i])
			jv := reflect.ValueOf(slice[j])

			return fd.Set(iv.Field(i)).Interface().(Sorter).Less(fd.Set(jv.Field(i)).Interface().(Sorter))
		})
	} else {
		panic("field not found")
	}
}

// Main function to demonstrate sorting
func main() {
	people := []Person{
		{Name: "Alice", Age: 30, Email: "alice@example.com"},
		{Name: "Bob", Age: 25, Email: "bob@example.com"},
		{Name: "Charlie", Age: 35, Email: "charlie@example.com"},
		{Name: "Alice", Age: 28, Email: "alice2@example.com"},
	}

	fmt.Println("Original people:")
	for _, person := range people {
		fmt.Printf("%+v\n", person)
	}

	// Sort by Name
	SortByField(sort.Slice(people, func(i, j int) bool {
		return people[i].Less(people[j])
	}), "Name")

	fmt.Println("\nSorted by Name:")
	for _, person := range people {
		fmt.Printf("%+v\n", person)
	}

	// Sort by Age
	SortByField(sort.Slice(people, func(i, j int) bool {
		return people[i].Less(people[j])
	}), "Age")

	fmt.Println("\nSorted by Age:")
	for _, person := range people {
		fmt.Printf("%+v\n", person)
	}

	// Sort by Email
	SortByField(sort.Slice(people, func(i, j int) bool {
		return people[i].Less(people[j])
	}), "Email")

	fmt.Println("\nSorted by Email:")
	for _, person := range people {
		fmt.Printf("%+v\n", person)
	}
}