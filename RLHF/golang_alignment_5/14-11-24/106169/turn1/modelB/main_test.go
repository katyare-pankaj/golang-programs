package main

import (
	"testing"
)

func TestScrubName(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"Alice!", "Alice"},
		{"Bob 2021", "Bob"},
		{"Charlie&David", "Charlie David"},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got := scrubName(tt.input)
			if got != tt.want {
				t.Errorf("scrubName(%q) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}
