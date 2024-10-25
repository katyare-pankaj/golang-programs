package main

import (
	"bytes"
	"compress/gzip"
	"compress/zlib"
	"fmt"
	"io/ioutil"
	"log"
	"testing"
	"time"
)

func main() {
	// Original data
	data := []byte("This is a sample data string that will be compressed using different methods to compare their effects on processing time in real-time applications.")

	fmt.Println("Original data size:", len(data))

	// Compare Gzip and Zlib compression
	testCompressionMethods(data, compressDataGzip, "Gzip")
	testCompressionMethods(data, compressDataZlib, "Zlib")
}

func testCompressionMethods(data []byte, compressionFunc func([]byte) []byte, methodName string) {
	start := time.Now()
	compressedData := compressionFunc(data)
	elapsed := time.Since(start)

	fmt.Printf("%s Compression Time: %s\n", methodName, elapsed)
	fmt.Printf("%s Compressed data size: %d\n", methodName, len(compressedData))

	// Decompress and verify
	decompressedData := decompressData(compressedData)

	if bytes.Equal(data, decompressedData) {
		fmt.Println("Decompression successful: Data matches original")
	} else {
		fmt.Println("Decompression failed: Data mismatch")
	}
}

func compressDataGzip(data []byte) []byte {
	var buf bytes.Buffer
	zw := gzip.NewWriter(&buf)
	if _, err := zw.Write(data); err != nil {
		log.Fatal(err)
	}
	if err := zw.Close(); err != nil {
		log.Fatal(err)
	}
	return buf.Bytes()
}

func compressDataZlib(data []byte) []byte {
	var buf bytes.Buffer
	w := zlib.NewWriter(&buf)
	if _, err := w.Write(data); err != nil {
		log.Fatal(err)
	}
	if err := w.Close(); err != nil {
		log.Fatal(err)
	}
	return buf.Bytes()
}

func decompressData(compressedData []byte) []byte {
	gr, err := gzip.NewReader(bytes.NewBuffer(compressedData))
	if err != nil {
		log.Fatal(err)
	}
	defer gr.Close()
	data, err := ioutil.ReadAll(gr)
	if err != nil {
		log.Fatal(err)
	}
	return data
}

func BenchmarkGzipCompression(b *testing.B) {
	data := []byte("This is some sample data for benchmarking")
	for n := 0; n < b.N; n++ {
		compressDataGzip(data)
	}
}

func BenchmarkZlibCompression(b *testing.B) {
	data := []byte("This is some sample data for benchmarking")
	for n := 0; n < b.N; n++ {
		compressDataZlib(data)
	}
}
