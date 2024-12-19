package main

import (
	"fmt"
	"reflect"
)

func describe(v interface{}) {
	val := reflect.ValueOf(v)
	typ := reflect.TypeOf(v)

	fmt.Printf("Value: %v\n", v)
	fmt.Printf("Type: %v\n", typ)
	fmt.Printf("Kind: %v\n", val.Kind())
	fmt.Println()
}

func main() {
	describe(42)
	describe(3.14)
	describe("hello")
	describe([]int{1, 2, 3})
	describe(struct{ Name string }{"Alice"})
}
