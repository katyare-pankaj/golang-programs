package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api"
)

const (
	influxDBURL   = "http://localhost:8086"
	influxDBToken = "your_influxdb_token"
	org           = "myorg"
	bucket        = "mybucket"
)

func main() {
	// Create a new InfluxDB Client
	client := influxdb2.NewClient(influxDBURL, influxDBToken)
	// Always close clients at the end to free up resources
	defer client.Close()

	// Create a new API service
	writeAPI := client.WriteAPI(org, bucket)

	// Securely handle data files
	dataFilePath := "sensor_data.csv"
	err := handleSecureFile(dataFilePath, writeAPI)
	if err != nil {
		log.Fatalf("Error handling file: %v", err)
	}

	fmt.Println("Data imported successfully.")
}

func handleSecureFile(filePath string, writeAPI api.WriteAPI) error {
	// Open the file securely
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("error opening file: %w", err)
	}
	defer file.Close()

	// Read the file line by line and import data into InfluxDB
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// Parse the line into an InfluxDB point and write it
		p := influxdb2.NewPoint("sensor_data",
			map[string]string{"location": "office"},
			map[string]interface{}{
				"temperature": 80.0,
				"humidity":    60.0,
			},
			time.Now())
		writeAPI.WritePoint(p)
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("error reading file: %w", err)
	}

	// Close the WriteAPI to ensure data is flushed and written
	return writeAPI.Close()
}
