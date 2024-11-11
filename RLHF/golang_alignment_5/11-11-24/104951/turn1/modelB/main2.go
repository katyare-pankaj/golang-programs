// Example of log data compression in Go using gzip
package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
)

func compressData(data []byte) ([]byte, error) {
	var buf bytes.Buffer
	gz := gzip.NewWriter(&buf)
	if _, err := gz.Write(data); err != nil {
		return nil, err
	}
	if err := gz.Close(); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func main() {
	logData := []byte("This is some log data")
	compressedData, err := compressData(logData)
	if err != nil {
		fmt.Println("Error compressing data:", err)
		return
	}
	fmt.Println("Compressed data size:", len(compressedData))
	// Store the 'compressedData' in CDN
}
