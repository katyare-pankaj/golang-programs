package main

import (
	"bytes"
	"compress/gzip"
	"io/ioutil"
	"math/rand"
	"testing"

	"github.com/andybalholm/brotli"
)

func generateRandomData(size int) []byte {
	data := make([]byte, size)
	rand.Read(data)
	return data
}

func compressDataGzip(data []byte) []byte {
	var buf bytes.Buffer
	zw := gzip.NewWriter(&buf)
	if _, err := zw.Write(data); err != nil {
		panic(err)
	}
	if err := zw.Close(); err != nil {
		panic(err)
	}
	return buf.Bytes()
}

func decompressDataGzip(compressedData []byte) []byte {
	gr, err := gzip.NewReader(bytes.NewBuffer(compressedData))
	if err != nil {
		panic(err)
	}
	defer gr.Close()
	data, err := ioutil.ReadAll(gr)
	if err != nil {
		panic(err)
	}
	return data
}

func compressDataBrotli(data []byte) []byte {
	var buf bytes.Buffer
	w := brotli.NewWriter(&buf)
	if _, err := w.Write(data); err != nil {
		panic(err)
	}
	if err := w.Close(); err != nil {
		panic(err)
	}
	return buf.Bytes()
}

func decompressDataBrotli(compressedData []byte) []byte {
	r := brotli.NewReader(bytes.NewBuffer(compressedData))
	data, err := ioutil.ReadAll(r)
	if err != nil {
		panic(err)
	}
	return data
}

func benchmarkCompression(b *testing.B, compressionFn func([]byte) []byte) {
	data := generateRandomData(1024 * 1024) // 1MB data
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		compressedData := compressionFn(data)
		decompressDataGzip(compressedData)
	}
}

func main() {
	testing.Benchmark(BenchmarkCompressionGzip)
	testing.Benchmark(BenchmarkCompressionBrotli)
}

func BenchmarkCompressionGzip(b *testing.B) {
	benchmarkCompression(b, compressDataGzip)
}

func BenchmarkCompressionBrotli(b *testing.B) {
	benchmarkCompression(b, compressDataBrotli)
}
