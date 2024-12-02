package logger_test

import (
	"bytes"
	"go-programs/RLHF/golang_random/2-12-24/389144/turn1/modelA/logger"
	"io/ioutil"

	"testing"
)

type testOut struct {
	bytes.Buffer
}

func (t *testOut) Write(b []byte) (int, error) {
	return t.Buffer.Write(b)
}

func TestLogMessage(t *testing.T) {
	// Create a mock writer
	testWriter := new(testOut)

	// Log a message
	logger.LogMessage(testWriter, "Hello, this is a test message!")

	// Read and validate the output
	output, err := ioutil.ReadAll(testWriter)
	if err != nil {
		t.Fatal(err)
	}

	expectedOutput := "Hello, this is a test message!\n"
	if string(output) != expectedOutput {
		t.Errorf("expected output: %v, got: %v", expectedOutput, string(output))
	}
}
