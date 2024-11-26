package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRegisterHandler(t *testing.T) {
	testCases := []struct {
		name     string
		body     []byte
		expected string
		status   int
	}{
		{
			name:     "Valid user registration",
			body:     []byte(`{"username":"testuser","password":"testpass"}`),
			expected: `{"message":"User registered successfully"}`,
			status:   http.StatusOK,
		},
		{
			name:     "Invalid request method",
			body:     []byte(`{"username":"testuser","password":"testpass"}`),
			expected: "Method not allowed",
			status:   http.StatusMethodNotAllowed,
		},
		{
			name:     "Invalid request body",
			body:     []byte(`{"username":"testuser"}`),
			expected: "Invalid request body",
			status:   http.StatusBadRequest,
		},
		{
			name:     "Invalid username or password",
			body:     []byte(`{"username":"","password":"testpass"}`),
			expected: "Invalid username or password",
			status:   http.StatusBadRequest,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			req, err := http.NewRequest("POST", "/register", bytes.NewBuffer(tc.body))
			if err != nil {
				t.Fatal(err)
			}

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(registerHandler)
			handler.ServeHTTP(rr, req)

			if status := rr.Code; status != tc.status {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, tc.status)
			}

			expected := tc.expected
			if rr.Body.String() != expected {
				t.Errorf("handler returned unexpected body: got %v want %v",
					rr.Body.String(), expected)
			}
		})
	}
}
