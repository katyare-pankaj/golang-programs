package main

import (
	"fmt"
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

// Function to compare two schemas and find the mismatches
func SchemaMatch(table1 Table, table2 Table) map[string]string {
	mismatches := make(map[string]string)

	// Compare table names
	if table1.Name != table2.Name {
		mismatches["Table Name"] = fmt.Sprintf("Mismatched: %s != %s", table1.Name, table2.Name)
	}

	// Find columns that exist in one table but not the other
	var columnNames1, columnNames2 []string
	for _, col := range table1.Columns {
		columnNames1 = append(columnNames1, col.Name)
	}
	for _, col := range table2.Columns {
		columnNames2 = append(columnNames2, col.Name)
	}

	diff := make(map[string]int)
	for _, name := range columnNames1 {
		diff[name] = 1
	}
	for _, name := range columnNames2 {
		diff[name] = diff[name] - 1
	}

	for name, count := range diff {
		if count != 0 {
			mismatches["Columns Exist"] = fmt.Sprintf("Mismatched: %s found in %s table only", name, tableName(count))
		}
	}

	// Compare columns that exist in both tables
	for _, col1 := range table1.Columns {
		for _, col2 := range table2.Columns {
			if col1.Name == col2.Name {
				if col1.DataType != col2.DataType {
					mismatches["Data Types"] = fmt.Sprintf("Mismatched: %s -> %s != %s", col1.Name, col1.DataType, col2.DataType)
				}
			}
		}
	}

	return mismatches
}

func tableName(count int) string {
	if count > 0 {
		return "First"
	}
	return "Second"
}

func main() {
	// Sample data
	table1 := Table{
		Name: "Orders",
		Columns: []Column{
			{Name: "order_id", DataType: "INT"},
			{Name: "customer_id", DataType: "INT"},
			{Name: "order_date", DataType: "DATE"},
		},
	}

	table2 := Table{
		Name: "Orders", // Same name as table1
		Columns: []Column{
			{Name: "order_id", DataType: "INTEGER"}, // Mismatched data type
			{Name: "customer_id", DataType: "INT"},
			{Name: "order_quantity", DataType: "INT"}, // Column missing in table1
		},
	}

	mismatches := SchemaMatch(table1, table2)
	for key, value := range mismatches {
		fmt.Printf("%s: %s\n", key, value)
	}
}
