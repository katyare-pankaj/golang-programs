package main

import (
	"fmt"
)

// DataMetadata represents metadata about a data item
type DataMetadata struct {
	Name          string   `json:"name"`
	Description   string   `json:"description"`
	CreatedAt     string   `json:"created_at"`
	LastUpdatedAt string   `json:"last_updated_at"`
	Owner         string   `json:"owner"`
	Size          int64    `json:"size"`
	Format        string   `json:"format"`
	Tags          []string `json:"tags"`
}

func main() {
	// Create a sample DataMetadata instance
	metadata := DataMetadata{
		Name:          "Employee Records",
		Description:   "Contains employee information such as name, age, and department.",
		CreatedAt:     "2023-07-28T10:00:00Z",
		LastUpdatedAt: "2023-07-28T10:00:00Z",
		Owner:         "john.doe@example.com",
		Size:          102400,
		Format:        "CSV",
		Tags:          []string{"employee", "data", "human_resources"},
	}

	// Display the metadata information
	fmt.Println("Data Metadata:")
	fmt.Printf("Name: %s\n", metadata.Name)
	fmt.Printf("Description: %s\n", metadata.Description)
	fmt.Printf("Created At: %s\n", metadata.CreatedAt)
	fmt.Printf("Last Updated At: %s\n", metadata.LastUpdatedAt)
	fmt.Printf("Owner: %s\n", metadata.Owner)
	fmt.Printf("Size: %d bytes\n", metadata.Size)
	fmt.Printf("Format: %s\n", metadata.Format)
	fmt.Printf("Tags: %s\n", metadata.Tags)
}
