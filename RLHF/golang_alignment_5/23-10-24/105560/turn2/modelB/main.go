package main

import (
	"fmt"
)

const (
	// Fuzzy matching threshold
	fuzzyMatchThreshold = 0.6
)

// Define a Table struct to represent a table schema
type Table struct {
	Name    string
	Columns []Column
}

// Define a Column struct to represent a column in a table schema
type Column struct {
	Name     string
	DataType string
}

// Define a PredefinedSchema struct to represent the predefined schema metadata
type PredefinedSchema struct {
	Tables map[string]TableMetadata
}

// Define a TableMetadata struct to represent the metadata of a table
type TableMetadata struct {
	Columns map[string]ColumnMetadata
}

// Define a ColumnMetadata struct to represent the metadata of a column
type ColumnMetadata struct {
	DataType    string
	Nullable    bool
	Constraints []string
}

// Function to match schemas using predefined schema metadata
func MatchSchemas(predefinedSchema PredefinedSchema, actualSchema []Table) map[string]string {
	mismatches := make(map[string]string)

	// Check if the number of tables in the actual schema matches the predefined schema
	if len(actualSchema) != len(predefinedSchema.Tables) {
		mismatches["Table Count"] = fmt.Sprintf("Mismatched: %d tables in predefined schema, but %d tables in actual schema", len(predefinedSchema.Tables), len(actualSchema))
	}

	// Compare each table in the actual schema with the predefined schema
	for _, actualTable := range actualSchema {
		tableName := actualTable.Name
		if _, exists := predefinedSchema.Tables[tableName]; !exists {
			mismatches[tableName] = "Table not found in predefined schema"
			continue
		}

		// Compare columns in the actual table with the predefined table
		for _, actualColumn := range actualTable.Columns {
			columnName := actualColumn.Name
			if _, exists := predefinedSchema.Tables[tableName].Columns[columnName]; !exists {
				mismatches[fmt.Sprintf("%s.%s", tableName, columnName)] = "Column not found in predefined schema"
				continue
			}

			// Compare data types, nullability, and constraints
			predefinedColumn := predefinedSchema.Tables[tableName].Columns[columnName]
			if actualColumn.DataType != predefinedColumn.DataType {
				mismatches[fmt.Sprintf("%s.%s", tableName, columnName)] = fmt.Sprintf("Mismatched data type: %s != %s", actualColumn.DataType, predefinedColumn.DataType)
			}
			if actualColumn.Name != predefinedColumn.DataType {
				mismatches[fmt.Sprintf("%s.%s", tableName, columnName)] = fmt.Sprintf("Mismatched nullability: %t != %t", actualColumn, predefinedColumn.Nullable)
			}

			// Compare constraints (not implemented in this example)
		}
	}

	return mismatches
}

func main() {
	// Sample predefined schema metadata
	predefinedSchema := PredefinedSchema{
		Tables: map[string]TableMetadata{
			"Orders": {
				Columns: map[string]ColumnMetadata{
					"order_id":     {DataType: "INT", Nullable: false, Constraints: []string{"PRIMARY KEY"}},
					"customer_id":  {DataType: "INT", Nullable: false},
					"order_date":   {DataType: "DATE", Nullable: false},
					"order_status": {DataType: "VARCHAR", Nullable: false, Constraints: []string{"CHECK (order_status IN ('PENDING', 'COMPLETED', 'CANCELLED'))"}},
				},
			},
		},
	}

	// Sample actual schema
	actualSchema := []Table{
		{
			Name: "Orders",
			Columns: []Column{
				{Name: "order_id", DataType: "INT"},
				{Name: "customer_id", DataType: "INT"},
				{Name: "order_date", DataType: "DATE"},
				{Name: "order_status", DataType: "VARCHAR"},
			},
		},
	}

	mismatches := MatchSchemas(predefinedSchema, actualSchema)
	for key, value := range mismatches {
		fmt.Printf("%s: %s\n", key, value)
	}
}
