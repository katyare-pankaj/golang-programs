package main

import (
	"fmt"
	"reflect"
	"sync"
)

// CallbackFunc is an interface that defines the signature of a callback function.
type CallbackFunc interface {
	Call(args ...interface{})
}

// Callback is a struct that wraps a function and implements the CallbackFunc interface.
type Callback struct {
	fn   interface{}
	typ  reflect.Type
	args []reflect.Type
}

// Call invokes the callback function with the provided arguments.
func (cb *Callback) Call(args ...interface{}) {
	if len(args) != len(cb.args) {
		panic(fmt.Sprintf("expected %d arguments, got %d", len(cb.args), len(args)))
	}
	argValues := make([]reflect.Value, len(args))
	for i, arg := range args {
		argValues[i] = reflect.ValueOf(arg)
	}
	callFunc := reflect.ValueOf(cb.fn).MethodByName("Call")
	callFunc.Call(argValues)
}

// RegisterCallback adds a callback function to the callback map.
func RegisterCallback(name string, fn interface{}) {
	callbacksMutex.Lock()
	defer callbacksMutex.Unlock()

	// Get the type of the function
	fnType := reflect.TypeOf(fn)
	if fnType.Kind() != reflect.Func {
		panic("not a function")
	}

	// Extract the argument types
	args := make([]reflect.Type, fnType.NumIn())
	for i := 0; i < fnType.NumIn(); i++ {
		args[i] = fnType.In(i)
	}

	// Create and store the Callback struct
	callbacks[name] = &Callback{fn: fn, typ: fnType, args: args}
}

// ExecuteCallback invokes a registered callback with the specified arguments.
func ExecuteCallback(name string, args ...interface{}) {
	callbacksMutex.Lock()
	defer callbacksMutex.Unlock()

	if cb, ok := callbacks[name]; ok {
		cb.Call(args...)
	} else {
		panic(fmt.Sprintf("callback not found: %s", name))
	}
}

// Example callback functions
func exampleCallback1(message string) {
	fmt.Println("Example Callback 1:", message)
}

func exampleCallback2(a int, b string) {
	fmt.Println("Example Callback 2:", a, b)
}

var callbacks = make(map[string]*Callback)
var callbacksMutex sync.Mutex

func main() {
	RegisterCallback("cb1", exampleCallback1)
	RegisterCallback("cb2", exampleCallback2)

	ExecuteCallback("cb1", "Hello, world!") // Output: Example Callback 1: Hello, world!
	ExecuteCallback("cb2", 42, "Life")      // Output: Example Callback 2: 42 Life

	// This will panic because "cb3" is not registered
	// ExecuteCallback("cb3", "Oops!")
}
