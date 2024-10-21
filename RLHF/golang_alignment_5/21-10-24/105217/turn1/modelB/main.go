package main

import "fmt"

func sumList(list []int) int {
	total := 0
	for _, value := range list {
		total += value
	}
	return total
}

func main() {
	myList := []int{1, 2, 3, 4, 5}
	result := sumList(myList)
	fmt.Println("Sum of the list is:", result)
}
