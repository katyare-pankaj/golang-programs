package main

import (
	"fmt"
	"reflect"
	"testing"
)

// Function to be tested with range loop
func sumElements(v interface{}) int {
	total := 0
	val := reflect.ValueOf(v)
	switch val.Kind() {
	case reflect.Array, reflect.Slice:
		for i := 0; i < val.Len(); i++ {
			elem := val.Index(i)
			switch elem.Kind() {
			case reflect.Int:
				total += int(elem.Int())
			case reflect.Float64:
				total += int(elem.Float())
			default:
				panic(fmt.Sprintf("Unsupported element type: %T", elem.Interface()))
			}
		}
	default:
		panic(fmt.Sprintf("Unsupported value type: %T", v))
	}
	return total
}

// TestSumElements function uses the reflect package to dynamically test the sumElements function
func TestSumElements(t *testing.T) {
	testCases := []struct {
		name  string
		input interface{}
		want  int
	}{
		{
			name:  "Sum of integers",
			input: []int{1, 2, 3, 4},
			want:  10,
		},
		{
			name:  "Sum of floats",
			input: []float64{1.5, 2.5, 3.5},
			want:  9,
		},
		{
			name:  "Sum of mixed integers and floats",
			input: []interface{}{1, 2.0, 3, 4.5},
			want:  10,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := sumElements(tc.input)
			if got != tc.want {
				t.Errorf("sumElements(%v) = %v, want %v", tc.input, got, tc.want)
			}
		})
	}
}
