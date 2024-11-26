package main

import (
	"time"
)

// Metadata represents the metadata of a dataset
type Metadata struct {
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Source      string     `json:"source"`
	Schema      string     `json:"schema"`
	CreatedAt   time.Time  `json:"createdAt"`
	UpdatedAt   time.Time  `json:"updatedAt"`
	Governance  Governance `json:"governance"`
}

// Governance represents the governance rules associated with a dataset
type Governance struct {
	Owner         string      `json:"owner"`
	AccessControl []string    `json:"accessControl"`
	Retention     int         `json:"retention"` // Retention in days
	Compliance    []string    `json:"compliance"`
	DataQuality   DataQuality `json:"dataQuality"`
}

// DataQuality represents the data quality rules
type DataQuality struct {
	Accuracy     float64 `json:"accuracy"`
	Completeness float64 `json:"completeness"`
	Consistency  float64 `json:"consistency"`
}

func main() {
	// Example usage
	dataMetadata := Metadata{
		Name:        "CustomerData",
		Description: "Data related to customers",
		Source:      "CRM System",
		Schema:      "customer_id, first_name, last_name, email",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		Governance: Governance{
			Owner:         "data_team@example.com",
			AccessControl: []string{"read", "write"},
			Retention:     365,
			Compliance:    []string{"GDPR", "CCPA"},
			DataQuality: DataQuality{
				Accuracy:     0.95,
				Completeness: 0.98,
				Consistency:  0.99,
			},
		},
	}

	// Print metadata
	printMetadata(dataMetadata)
}

func printMetadata(m Metadata) {
	print("Metadata:")
	print(m)
}
