package main

import (
	"fmt"
	"syscall/js"
)

func main() {
	js.Global().Set("handleInput", js.FuncOf(handleInput))
	<-make(chan struct{}, 0)
}

func handleInput(this js.Value, args []js.Value) interface{} {
	if len(args) != 1 {
		panic("Expected exactly one argument")
	}
	userInput := args[0].String()
	// Safely escape the user input using fmt.Sprintf
	escapedInput := fmt.Sprintf("%q", userInput)
	js.Global().Call("console.log", "Escaped user input:", escapedInput)
	return nil
}
