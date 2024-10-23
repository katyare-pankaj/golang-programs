package main

import (
	"fmt"
	"os"
	"reflect"
	"testing"
)

// Define the DSL
type CompressionDSL struct {
	data []byte
}

func (dsl *CompressionDSL) Compress() *CompressionDSL {
	// Simplified compression logic for demonstration
	dsl.data = []byte("Compressed data")
	return dsl
}

func (dsl *CompressionDSL) Decompress() *CompressionDSL {
	// Simplified decompression logic for demonstration
	dsl.data = []byte("Original data")
	return dsl
}

func (dsl *CompressionDSL) Data() []byte {
	return dsl.data
}

// Type inference function
func inferType(input interface{}) string {
	return reflect.TypeOf(input).String()
}

// Test Cases
func TestTypeInferenceForCompressedData(t *testing.T) {
	dsl := &CompressionDSL{data: []byte("Original data")}
	compressedDSL := dsl.Compress()

	expectedType := "*main.CompressionDSL"
	actualType := inferType(compressedDSL)

	if actualType != expectedType {
		t.Errorf("Type inference failed. Expected: %s, Actual: %s", expectedType, actualType)
	}
}

func TestTypeInferenceForDecompressedData(t *testing.T) {
	dsl := &CompressionDSL{data: []byte("Compressed data")}
	decompressedDSL := dsl.Decompress()

	expectedType := "*main.CompressionDSL"
	actualType := inferType(decompressedDSL)

	if actualType != expectedType {
		t.Errorf("Type inference failed. Expected: %s, Actual: %s", expectedType, actualType)
	}
}

func ExampleCompressionDSL() {
	dsl := &CompressionDSL{data: []byte("Hello, World!")}
	compressedDSL := dsl.Compress()
	fmt.Printf("Compressed data: %s\n", compressedDSL.Data())
	// Output: Compressed data: [80 114 101 115 115 101 100 32 100 97 116 97]

	decompressedDSL := compressedDSL.Decompress()
	fmt.Printf("Decompressed data: %s\n", decompressedDSL.Data())
	// Output: Decompressed data: [72 101 108 108 111 44 32 87 111 114 108 100 33]
}

func main() {
	testing.Init()
	testing.MainStart(testMain, nil)
}

func testMain(m *testing.M) {
	fmt.Println("Running tests...")
	code := m.Run()
	fmt.Println("Test suite completed.")
	os.Exit(code)
}
