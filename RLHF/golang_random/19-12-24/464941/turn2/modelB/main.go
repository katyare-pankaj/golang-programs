package main

import (
	"fmt"
	"strconv"
)

func main() {
	unsafeStrings := []string{"1", "2", "three", "4", "5", "-6", "7.8", "9"}
	safeInts := make([]int, 0, len(unsafeStrings))

	for _, s := range unsafeStrings {
		if i, err := strconv.ParseInt(s, 10, 64); err == nil {
			safeInts = append(safeInts, int(i))
		} else {
			fmt.Println("Invalid input:", s)
		}
	}
	fmt.Println("Safe integers:", safeInts)
}
