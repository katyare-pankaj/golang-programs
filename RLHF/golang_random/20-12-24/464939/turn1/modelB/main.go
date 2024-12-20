package main

import (
	"fmt"
)

func main() {
	nestedMap := map[string]map[string]string{
		"animal": {
			"cat": "meow",
			"dog": "woof",
		},
		"human": {
			"male":   "hello",
			"female": "hi",
		},
	}

	for outerKey, innerMap := range nestedMap {
		fmt.Printf("%s: {\n", outerKey)
		for innerKey, value := range innerMap {
			fmt.Printf("  %s: %s,\n", innerKey, value)
		}
		fmt.Println("}")
	}
}
