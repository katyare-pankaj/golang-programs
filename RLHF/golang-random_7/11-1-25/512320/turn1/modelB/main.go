package main

import "fmt"

func main() {
	slice1 := []int{1, 2, 3}
	f1(slice1)
}

func f1(slice []int) {
	for _, value := range slice {
		f2(value)
	}
}

func f2(num int) {
	fmt.Println(num)
}
