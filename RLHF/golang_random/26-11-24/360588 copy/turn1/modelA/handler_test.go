package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHomeHandler(t *testing.T) {
	// Create a request to send to the handler
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	// Create a recorder to capture HTTP response
	recorder := httptest.NewRecorder()

	// Get the handler and route the request to it
	handler := http.HandlerFunc(HomeHandler)
	handler.ServeHTTP(recorder, req)

	// Check the status code is what we expect.
	if recorder.Code != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			recorder.Code, http.StatusOK)
	}

	// Check the response body is what we expect.
	if recorder.Body.String() != "Hello, World!\n" {
		t.Errorf("handler returned unexpected body: got %v want %v",
			recorder.Body.String(), "Hello, World!\n")
	}
}
