package main

import (
	"bytes"
	"compress/gzip"
	"compress/zlib"
	"fmt"
	"io/ioutil"
	"testing"
	"time"

	"github.com/klauspost/compress/snappy"
)

func compressData(data []byte, algo string) ([]byte, error) {
	switch algo {
	case "gzip":
		var b bytes.Buffer
		gz := gzip.NewWriter(&b)
		_, err := gz.Write(data)
		if err != nil {
			return nil, err
		}
		if err := gz.Close(); err != nil {
			return nil, err
		}
		return b.Bytes(), nil
	case "zlib":
		return zlib.Compress(data)
	case "snappy":
		return snappy.Encode(nil, data)
	default:
		return nil, fmt.Errorf("unknown compression algorithm: %s", algo)
	}
}
func decompressData(data []byte, algo string) ([]byte, error) {
	switch algo {
	case "gzip":
		gz, err := gzip.NewReader(bytes.NewReader(data))
		if err != nil {
			return nil, err
		}
		defer gz.Close()
		return ioutil.ReadAll(gz)
	case "zlib":
		return zlib.Decompress(data)
	case "snappy":
		return snappy.Decode(nil, data)
	default:
		return nil, fmt.Errorf("unknown compression algorithm: %s", algo)
	}
}

func TestCompressionAlgorithms(t *testing.T) {
	algorithms := []string{"gzip", "zlib", "snappy"}
	data := []byte("This is a sample test data. This is a sample test data. This is a sample test data. This is a sample test data. This is a sample test data.")

	for _, algo := range algorithms {
		t.Run(algo, func(t *testing.T) {
			compressedData, err := compressData(data, algo)
			if err != nil {
				t.Fatalf("Error compressing data with %s: %v", algo, err)
			}

			startTime := time.Now()
			decompressedData, err := decompressData(compressedData, algo)
			if err != nil {
				t.Fatalf("Error decompressing data with %s: %v", algo, err)
			}
			endTime := time.Now()

			if !bytes.Equal(data, decompressedData) {
				t.Fatalf("Decompressed data does not match the original data for %s", algo)
			}

			duration := endTime.Sub(startTime)
			t.Logf("Compression and Decompression time with %s: %v", algo, duration)
		})
	}
}

func main() {
	testing.Main()
}
