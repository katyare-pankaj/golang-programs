package main

import (
	"fmt"
	"sync"
)

type CallbackFunc func(args ...interface{})

var callbacks = make(map[string]CallbackFunc)
var callbacksMutex sync.Mutex

func RegisterCallback(name string, callback CallbackFunc) {
	callbacksMutex.Lock()
	defer callbacksMutex.Unlock()
	callbacks[name] = callback
}

func ExecuteCallback(name string, args ...interface{}) {
	callbacksMutex.Lock()
	defer callbacksMutex.Unlock()

	if callback, ok := callbacks[name]; ok {
		callback(args...)
	} else {
		fmt.Println("Callback not found:", name)
	}
}

func callback1(args ...interface{}) {
	fmt.Println("Callback 1:", args[0].(string))
}

func callback2(args ...interface{}) {
	fmt.Println("Callback 2:", args[0].(int), args[1].(string))
}

func main() {
	RegisterCallback("cb1", callback1)
	RegisterCallback("cb2", callback2)

	ExecuteCallback("cb1", "Hello!")
	ExecuteCallback("cb2", 42, "Gopher")

	ExecuteCallback("cb3", "Oops!") // This will result in a callback not found message
}
