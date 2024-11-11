package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"hash/fnv"
	"net/http"
	"strconv"

	"github.com/spf13/afero"
)

const (
	cdnURL      = "https://example-cdn.com/logs/"
	numShards   = 10
	logFilePath = "logfile.log"
)

func uploadLogDataToCDN(shardNumber int, logData []byte) error {
	url := cdnURL + strconv.Itoa(shardNumber)

	// Compress the log data before uploading
	var compressedData bytes.Buffer
	gz := gzip.NewWriter(&compressedData)
	if _, err := gz.Write(logData); err != nil {
		return fmt.Errorf("error compressing log data: %w", err)
	}
	if err := gz.Close(); err != nil {
		return fmt.Errorf("error closing gzip writer: %w", err)
	}

	req, err := http.NewRequest("PUT", url, &compressedData)
	if err != nil {
		return fmt.Errorf("error creating HTTP request: %w", err)
	}
	req.Header.Set("Content-Encoding", "gzip")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("error uploading log data to CDN: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("upload failed with status: %s", resp.Status)
	}

	return nil
}

func main() {
	// Read log data from file
	fs := afero.NewOsFs()
	logData, err := afero.ReadFile(fs, logFilePath)
	if err != nil {
		fmt.Println("Error reading log file:", err)
		return
	}

	// Determine the shard number based on a key (e.g., log file name or timestamp)
	// For simplicity, using the log file name as the key
	shardNumber := getShardNumber(logFilePath, numShards)

	// Upload the log data to the CDN
	if err := uploadLogDataToCDN(shardNumber, logData); err != nil {
		fmt.Println("Error uploading log data to CDN:", err)
		return
	}

	fmt.Println("Log data uploaded to CDN shard:", shardNumber)
}
func getShardNumber(key string, numShards int) int {
	// Simple hashing method for sharding
	h := fnv.New32a()
	h.Write([]byte(key))
	return int(h.Sum32() % uint32(numShards))
}
