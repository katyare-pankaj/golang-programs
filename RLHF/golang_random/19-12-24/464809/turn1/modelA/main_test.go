package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestPrintElements(t *testing.T) {
	testCases := []struct {
		name   string
		input  interface{}
		expect []string
	}{
		{
			name:   "Slice",
			input:  []int{1, 2, 3},
			expect: []string{"1", "2", "3"},
		},
		{
			name:   "Map",
			input:  map[string]int{"a": 1, "b": 2, "c": 3},
			expect: []string{"a=1", "b=2", "c=3"},
		},
		{
			name:   "Array",
			input:  [3]int{4, 5, 6},
			expect: []string{"4", "5", "6"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rv := reflect.ValueOf(tc.input)

			if rv.Kind() != reflect.Slice && rv.Kind() != reflect.Map && rv.Kind() != reflect.Array {
				t.Fatalf("input is not an iterable type: %v", rv.Kind())
			}

			var output []string
			for _, elem := range rv {
				// Check the kind to handle slice elements vs. map key/values
				if rv.Kind() == reflect.Map {
					key, val := elem.Interface().(map[string]int)["a"]
					output = append(output, fmt.Sprintf("a=%v", key))
				} else {
					// For slice and array
					output = append(output, fmt.Sprintf("%v", elem.Interface()))
				}
			}

			if !reflect.DeepEqual(output, tc.expect) {
				t.Errorf("expected %v, got %v", tc.expect, output)
			}
		})
	}
}
