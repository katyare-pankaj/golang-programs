package main

import (
	"fmt"
	"reflect"
	"time"
)

type ResourceIntensiveOps struct{}

func (r *ResourceIntensiveOps) Op1(param int) (string, error) {
	fmt.Printf("Processing with param: %d\n", param)
	time.Sleep(500 * time.Millisecond)

	if param%2 == 0 {
		return "", fmt.Errorf("even number error for Op1: %d", param)
	}
	return fmt.Sprintf("Success with Op1: %d", param), nil
}

func invokeMethod(obj interface{}, methodName string, params ...interface{}) (result []interface{}, err error) {
	v := reflect.ValueOf(obj)
	m := v.MethodByName(methodName)
	if !m.IsValid() {
		err = fmt.Errorf("no such method: %s", methodName)
		return
	}

	in := make([]reflect.Value, len(params))
	for i, param := range params {
		in[i] = reflect.ValueOf(param)
	}

	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("panic: %v", r)
		}
	}()
	out := m.Call(in)
	result = make([]interface{}, len(out))
	for i, o := range out {
		result[i] = o.Interface()
	}
	return
}

func main() {
	ops := &ResourceIntensiveOps{}
	// Using defer to ensure resource cleanup regardless of success or failure
	defer fmt.Println("Resources cleaned up.")

	result, err := invokeMethod(ops, "Op1", 4)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result[0])
	}
}
