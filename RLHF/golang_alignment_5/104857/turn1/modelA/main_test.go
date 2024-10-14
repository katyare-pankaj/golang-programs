package modelA

import (
	"reflect"
	"testing"
)

// BDD Test using Go's testing package
func TestSynchronization(t *testing.T) {
	type testCase struct {
		name     string
		source   []map[string]interface{}
		target   []map[string]interface{}
		expected []map[string]interface{}
	}

	testCases := []testCase{
		{
			name: "Simple synchronization",
			source: []map[string]interface{}{
				{"id": 1, "name": "Alice"},
			},
			target: []map[string]interface{}{
				{"id": 1, "name": "Alice"},
				{"id": 2, "name": "Bob"},
			},
			expected: []map[string]interface{}{
				{"id": 2, "name": "Bob"},
			},
		},
		// Add more test cases as needed
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			synchronizer := Synchronizer{}
			actual, err := synchronizer.SynchronizeData(tc.source, tc.target)
			if err != nil {
				t.Fatalf("Synchronization failed: %v", err)
			}

			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("Expected: %v, Actual: %v", tc.expected, actual)
			}
		})
	}
}
