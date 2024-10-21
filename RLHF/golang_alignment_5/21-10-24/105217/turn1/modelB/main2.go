package main

import "fmt"

func sumList(list []int) int {
	total := 0
	if len(list) > 0 { // Check if the list is not empty
		for _, value := range list {
			total += value
		}
	}
	return total
}

func main() {
	myList1 := []int{1, 2, 3, 4, 5}
	result1 := sumList(myList1)
	fmt.Println("Sum of the list is:", result1)

	myList2 := []int{} // Empty list
	result2 := sumList(myList2)
	fmt.Println("Sum of the empty list is:", result2) // Output will be 0
}
