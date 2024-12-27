package main

import "fmt"

type Callback func(...interface{})

var callbacks = make(map[string]Callback)

func RegisterCallback(name string, cb Callback) {
	callbacks[name] = cb
}

func ExecuteCallback(name string, args ...interface{}) {
	if cb, ok := callbacks[name]; ok {
		cb(args...)
	} else {
		fmt.Println("Callback not found:", name)
	}
}

func sayHello(args ...interface{}) {
	name := args[0].(string)
	fmt.Println("Hello,", name)
}

func main() {
	RegisterCallback("hello", sayHello)
	ExecuteCallback("hello", "Alice")
	ExecuteCallback("world", "Bob") // This will print "Callback not found: world"
}
