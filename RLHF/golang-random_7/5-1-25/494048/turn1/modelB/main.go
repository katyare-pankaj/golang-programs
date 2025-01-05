package main

import (
	"fmt"
)

func main() {
	fruits := []string{"apple", "banana", "orange"}
	for index, fruit := range fruits {
		fmt.Println("Index:", index, ", Fruit:", fruit)
	}
}
