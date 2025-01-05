package main

import (
	"bytes"
	"fmt"
	"math/rand"
	"strconv"
)

func simpleStringFormat() string {
	name := "Alice"
	age := 25
	return "Hello, my name is " + name + ". I am " + strconv.Itoa(age) + " years old."
}

func largeScaleStringFormat() string {
	var result string
	for _, name := range names {
		age := rand.Intn(100)
		result += fmt.Sprintf("%s is %d years old.\n", name, age)
	}
	return result
}

func optimizedLargeScaleStringFormat() string {
	var buf bytes.Buffer
	for _, name := range names {
		age := rand.Intn(100)
		fmt.Fprintf(&buf, "%s is %d years old.\n", name, age)
	}
	return buf.String()
}

func complexStringFormat() string {
	return fmt.Sprintf("%10s: %03d\n", "Score", 123)
}

func main() {
	simpleStringFormat()
	largeScaleStringFormat()
	optimizedLargeScaleStringFormat()
	complexStringFormat()
}
