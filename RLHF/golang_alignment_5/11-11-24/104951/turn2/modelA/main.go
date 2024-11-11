package main

import (
	"compress/gzip"
	"fmt"
	"hash/fnv"
	"io"
	"math/rand"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/spf13/afero"
)

const (
	cdnURL        = "https://your-cdn-domain.com/logs/"
	numShards     = 10
	logDataDir    = "logs"
	compressedDir = "compressed"
)

func getShardNumber(key string, numShards int) int {
	h := fnv.New32a()
	h.Write([]byte(key))
	return int(h.Sum32() % uint32(numShards))
}

func compressFile(src, dest string) error {
	f, err := os.Open(src)
	if err != nil {
		return err
	}
	defer f.Close()

	d, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer d.Close()

	gz := gzip.NewWriter(d)
	defer gz.Close()

	_, err = io.Copy(gz, f)
	return err
}

func uploadLogDataToCDN(shardNumber int, compressedLogPath string) error {
	url := cdnURL + strconv.Itoa(shardNumber)
	// Implement HTTP PUT request to upload the compressed file to the CDN URL
	// For simplicity, we'll just print the URL in this example
	fmt.Println("Uploading compressed log to:", url)
	return nil
}

func main() {
	// Initialize random seed
	rand.Seed(time.Now().UnixNano())

	fs := afero.NewOsFs()

	// Generate sample log data files in logDataDir
	for i := 0; i < 10; i++ {
		logData := []byte(fmt.Sprintf("log entry %d: some random data\n", i))
		logFilePath := filepath.Join(logDataDir, fmt.Sprintf("log-%d.txt", i))
		if err := afero.WriteFile(fs, logFilePath, logData, 0644); err != nil {
			fmt.Println("Error creating log file:", err)
			return
		}
	}

	// Shard and compress log data
	err := fs.MkdirAll(compressedDir, 0755)
	if err != nil {
		fmt.Println("Error creating compressed dir:", err)
		return
	}

	logFiles, err := afero.ReadDir(fs, logDataDir)
	if err != nil {
		fmt.Println("Error reading log files:", err)
		return
	}

	for _, logFile := range logFiles {
		logFilePath := filepath.Join(logDataDir, logFile.Name())
		compressedFilePath := filepath.Join(compressedDir, logFile.Name()+".gz")

		if err := compressFile(logFilePath, compressedFilePath); err != nil {
			fmt.Println("Error compressing file:", err)
			return
		}

		// Determine the shard number based on the log file name
		shardNumber := getShardNumber(logFile.Name(), numShards)

		// Upload the compressed log data to the CDN
		if err := uploadLogDataToCDN(shardNumber, compressedFilePath); err != nil {
			fmt.Println("Error uploading log data to CDN:", err)
			return
		}
	}

	fmt.Println("Log data sharding and compression completed.")
}
