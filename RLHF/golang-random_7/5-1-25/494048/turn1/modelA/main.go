package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	slice := []int{1, 2, 3, 4, 5}
	for i, v := range slice {
		fmt.Println(i, v) // Prints index and value
	}

	mapData := map[string]int{"a": 1, "b": 2, "c": 3}
	for key, value := range mapData {
		fmt.Println(key, value) // Prints key and value
	}

	str := "Hello, World!"
	for _, char := range str {
		fmt.Print(char) // Prints each character
	}

	ch := make(chan int)
	go func() {
		ch <- 1
		ch <- 2
		close(ch)
	}()
	for value := range ch {
		fmt.Println(value) // Prints values from the channel
	}

	file, err := os.Open("file.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text()) // Prints each line from the file
	}
}
