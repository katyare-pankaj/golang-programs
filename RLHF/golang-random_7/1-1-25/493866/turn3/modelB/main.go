package main

import (
	"fmt"
	"log"
)

func main() {
	var slice1 []int = nil
	var slice2 = []int{1, 2, 3}
	var arr1 [3]int = nil
	var arr2 = [3]int{4, 5, 6}

	// Logging nil slice and array
	log.Println(fmt.Sprintf("Nil slice: %v", slice1))
	log.Println(fmt.Sprintf("Non-nil slice: %v", slice2))
	log.Println(fmt.Sprintf("Nil array: %v", arr1))
	log.Println(fmt.Sprintf("Non-nil array: %v", arr2))
}
