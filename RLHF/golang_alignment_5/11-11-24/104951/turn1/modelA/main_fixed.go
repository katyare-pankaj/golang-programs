package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
	"sync"
	"time"

	"github.com/akamai/AkamaiOPEN-edgegrid-golang/edgegrid"
)

const (
	cdnBaseURL = "https://example.cdn.com/archive/"
	dataDir    = "data_to_archive"
)

var (
	wg sync.WaitGroup
)

// Function to create a new authenticated HTTP request with EdgeGrid headers
func createAuthenticatedRequest(config edgegrid.Config, method, url string, body []byte) (*http.Request, error) {
	req, err := http.NewRequest(method, url, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}

	// Add EdgeGrid authentication headers to the request
	req = edgegrid.AddRequestHeader(config, req)
	return req, nil
}

func archiveDataConcurrently(config edgegrid.Config, dataFiles []string) {
	for _, filePath := range dataFiles {
		wg.Add(1)
		go archiveData(config, filePath)
	}
	wg.Wait()
}

func archiveData(config edgegrid.Config, filePath string) {
	defer wg.Done()

	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Printf("Error reading file %s: %v", filePath, err)
		return
	}

	baseName := filepath.Base(filePath)
	url := cdnBaseURL + baseName

	req, err := createAuthenticatedRequest(config, "PUT", url, data)
	if err != nil {
		log.Printf("Error creating request for %s: %v", filePath, err)
		return
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Error archiving %s: %v", filePath, err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("Failed to archive %s: %s", filePath, resp.Status)
	} else {
		log.Printf("Archived %s successfully", filePath)
	}
}

func main() {
	// Initialize Akamai EdgeGrid client configuration
	config := edgegrid.Config{
		ClientToken:  "YOUR_CLIENT_TOKEN",
		ClientSecret: "YOUR_CLIENT_SECRET",
		AccessToken:  "YOUR_ACCESS_TOKEN",
		Host:         "YOUR_HOST",
	}

	// Get list of data files to archive
	dataFiles, err := filepath.Glob(filepath.Join(dataDir, "*"))
	if err != nil {
		log.Fatalf("Error finding data files: %v", err)
	}

	startTime := time.Now()
	log.Println("Archiving data concurrently...")
	archiveDataConcurrently(config, dataFiles)
	log.Println("Archiving completed in", time.Since(startTime))
}
