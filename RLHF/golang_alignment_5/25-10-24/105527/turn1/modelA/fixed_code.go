package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io/ioutil"
	"log"
	"testing"
	"time"

	"github.com/andybalholm/brotli"
)

func main() {
	// Original data
	data := []byte("This is some sample data that needs to be compressed for real-time marketing and advertising applications.")

	fmt.Println("Original data size:", len(data))

	// Test Gzip compression
	start := time.Now()
	compressedDataGzip := compressDataGzip(data)
	elapsed := time.Since(start)
	fmt.Println("Gzip Compression Time:", elapsed)
	fmt.Println("Gzip Compressed data size:", len(compressedDataGzip))

	// Test Brotli compression
	start = time.Now()
	compressedDataBrotli := compressDataBrotli(data)
	elapsed = time.Since(start)
	fmt.Println("Brotli Compression Time:", elapsed)
	fmt.Println("Brotli Compressed data size:", len(compressedDataBrotli))

	// Decompress and verify
	decompressedDataGzip := decompressDataGzip(compressedDataGzip)
	decompressedDataBrotli := decompressDataBrotli(compressedDataBrotli)

	if bytes.Equal(data, decompressedDataGzip) && bytes.Equal(data, decompressedDataBrotli) {
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

func decompressDataGzip(compressedData []byte) []byte {
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

func compressDataBrotli(data []byte) []byte {
	var buf bytes.Buffer
	w := brotli.NewWriter(&buf)
	if _, err := w.Write(data); err != nil {
		log.Fatal(err)
	}
	if err := w.Close(); err != nil {
		log.Fatal(err)
	}
	return buf.Bytes()
}

func decompressDataBrotli(compressedData []byte) []byte {
	r := brotli.NewReader(bytes.NewBuffer(compressedData))
	data, err := ioutil.ReadAll(r)
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

func BenchmarkBrotliCompression(b *testing.B) {
	data := []byte("This is some sample data for benchmarking")
	for n := 0; n < b.N; n++ {
		compressDataBrotli(data)
	}
}
