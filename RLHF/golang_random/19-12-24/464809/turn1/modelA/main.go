package main

import "fmt"

func printElements(iter interface{}) {
	for _, elem := range iter {
		fmt.Println(elem)
	}
}

func main() {
	// Let's print elements of different types
	slice := []int{1, 2, 3}
	mapExample := map[string]int{"a": 1, "b": 2, "c": 3}
	array := [3]int{4, 5, 6}

	printElements(slice)
	printElements(mapExample)
	printElements(array)
}
