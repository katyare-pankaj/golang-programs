package main

import "fmt"

func sumArrayReadable(array []int) int {
	total := 0
	for _, value := range array {
		total += value
	}
	return total
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	result := sumArrayReadable(arr)
	fmt.Println("Sum:", result)
}
