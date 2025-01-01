package main

import (
	"fmt"
	"testing"
)

func TestNilPointerFormat(t *testing.T) {
	var p *int = nil
	expected := "nil"
	actual := fmt.Sprintf("%v", p)
	if actual != expected {
		t.Errorf("Expected '%s', got '%s'", expected, actual)
	}
}

func TestNilInterfaceFormat(t *testing.T) {
	var i interface{} = nil
	expected := "<nil>"
	actual := fmt.Sprintf("%v", i)
	if actual != expected {
		t.Errorf("Expected '%s', got '%s'", expected, actual)
	}
}

func TestNilStringFormat(t *testing.T) {
	var s *string = nil
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic, but none occurred")
		}
	}()
	fmt.Sprintf("%v", s) // This will panic
}

func TestStringFormat(t *testing.T) {
	var s string = "Hello, World!"
	expected := "Hello, World!"
	actual := fmt.Sprintf("%v", s)
	if actual != expected {
		t.Errorf("Expected '%s', got '%s'", expected, actual)
	}
}
