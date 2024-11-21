package main

import (
	"reflect"
	"sort"
	"testing"
)

func TestFactorizeNumber(t *testing.T) {
	tests := []struct {
		number int
		want   []int
	}{
		{100, []int{2, 2, 5, 5}},
		{36, []int{2, 2, 3, 3}},
		{25, []int{5, 5}},
		{1, []int{}},
		{121, []int{11, 11}},
	}

	for _, test := range tests {
		got := factorizeNumber(test.number)
		if !testEqual(got, test.want) {
			t.Errorf("factorizeNumber(%d) = %v, want %v", test.number, got, test.want)
		}
	}
}

func testEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	sortedA := make([]int, len(a))
	copy(sortedA, a)
	sortedB := make([]int, len(b))
	copy(sortedB, b)
	sort.Ints(sortedA)
	sort.Ints(sortedB)
	return reflect.DeepEqual(sortedA, sortedB)
}
