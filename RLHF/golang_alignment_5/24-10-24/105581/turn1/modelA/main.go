package main

import (
	"testing"
)

// Define a struct to represent the data being synced
type UserData struct {
	ID        int
	Email     string
	Name      string
	UpdatedAt string
}

// Function to simulate data syncing from source to destination
func syncData(sourceData []UserData, destinationData []UserData) []UserData {
	// Implement the data syncing logic here
	// For simplicity, we'll just append the new data to the destination
	var updatedDestinationData []UserData
	updatedDestinationData = append(updatedDestinationData, destinationData...)
	for _, data := range sourceData {
		updatedDestinationData = append(updatedDestinationData, data)
	}
	return updatedDestinationData
}

func TestSyncDataCompatibility(t *testing.T) {
	testCases := []struct {
		name        string
		source      []UserData
		destination []UserData
		expected    []UserData
	}{
		{
			name: "Basic sync",
			source: []UserData{
				{ID: 1, Email: "user1@example.com", Name: "User One"},
			},
			destination: []UserData{},
			expected: []UserData{
				{ID: 1, Email: "user1@example.com", Name: "User One"},
			},
		},
		{
			name: "Sync with existing data",
			source: []UserData{
				{ID: 2, Email: "user2@example.com", Name: "User Two"},
			},
			destination: []UserData{
				{ID: 1, Email: "user1@example.com", Name: "User One"},
			},
			expected: []UserData{
				{ID: 1, Email: "user1@example.com", Name: "User One"},
				{ID: 2, Email: "user2@example.com", Name: "User Two"},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := syncData(tc.source, tc.destination)
			if !areDataSlicesEqual(result, tc.expected) {
				t.Errorf("Expected %v, got %v", tc.expected, result)
			}
		})
	}
}

func areDataSlicesEqual(data1 []UserData, data2 []UserData) bool {
	if len(data1) != len(data2) {
		return false
	}

	for i := range data1 {
		if data1[i] != data2[i] {
			return false
		}
	}

	return true
}
