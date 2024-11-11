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

func archiveDataConcurrently(client *edgegrid.EdgeGridClient, dataFiles []string) {
	for _, filePath := range dataFiles {
		wg.Add(1)
		go archiveData(client, filePath)
	}
	wg.Wait()
}

func archiveData(client *edgegrid.EdgeGridClient, filePath string) {
	defer wg.Done()

	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Printf("Error reading file %s: %v", filePath, err)
		return
	}

	baseName := filepath.Base(filePath)
	url := cdnBaseURL + baseName

	req, err := http.NewRequest("PUT", url, bytes.NewReader(data))
	if err != nil {
		log.Printf("Error creating request for %s: %v", filePath, err)
		return
	}

	// Set any required headers for authentication or other purposes
	// req.Header.Set("Authorization", "Bearer YOUR_ACCESS_TOKEN")

	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Error archiving %s: %v", filePath, err)
		return
	}

	if resp.StatusCode != http.StatusOK {
		log.Printf("Failed to archive %s: %s", filePath, resp.Status)
	} else {
		log.Printf("Archived %s successfully", filePath)
	}

	resp.Body.Close()
}

func main() {
	// Initialize Akamai EdgeGrid client (replace with your credentials)
	credentials := edgegrid.Config{
		ClientToken:  "YOUR_CLIENT_TOKEN",
		ClientSecret: "YOUR_CLIENT_SECRET",
		AccessToken:  "YOUR_ACCESS_TOKEN",
	}
	client, err := client_v1.New(credentials)
	if err != nil {
		log.Fatalf("Error creating EdgeGrid client: %v", err)
	}

	// Get list of data files to archive
	dataFiles, err := filepath.Glob(filepath.Join(dataDir, "*"))
	if err != nil {
		log.Fatalf("Error finding data files: %v", err)
	}

	startTime := time.Now()
	log.Println("Archiving data concurrently...")
	archiveDataConcurrently(client, dataFiles)
	log.Println("Archiving completed in", time.Since(startTime))
}
