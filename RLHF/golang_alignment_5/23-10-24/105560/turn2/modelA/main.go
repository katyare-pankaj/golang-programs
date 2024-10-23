package main

import (
   "fmt"
   "github.com/sahilm/fuzzy"
   "strings"
)

const (
   // Fuzzy matching threshold
   fuzzyMatchThreshold = 0.6
)

// Define a Table struct to represent a table schema
type Table struct {
   Name     string
   Columns []Column
}

// Define a Column struct to represent a column in a table schema
type Column struct {
   Name     string
   DataType string
}

// Function to perform fuzzy match between two strings
func fuzzyMatch(str1, str2 string) bool {
   score := fuzzy.MatchFuzzy(str1, str2, nil)
   return score >= fuzzyMatchThreshold
}

// Function to compare two schemas and find the mismatches
func SchemaMatch(table1 Table, table2 Table) map[string]string {
   mismatches := make(map[string]string)

   // Compare table names using fuzzy matching
   if !fuzzyMatch(table1.Name, table2.Name) {
   	mismatches["Table Name"] = fmt.Sprintf("Mismatched: %s != %s", table1.Name, table2.Name)
   }

   // Find columns that exist in one table but not the other using fuzzy matching on column names
   var columnNames1, columnNames2 []string
   for _, col := range table1.Columns {
   	columnNames1 = append(columnNames1, col.Name)
   }
   for _, col := range table2.Columns {
   	columnNames2 = append(columnNames2, col.Name)
   }

   // ... (rest of the code remains the same)

   // Compare columns that exist in both tables, allowing for data type variations
   for _, col1 := range table1.Columns {
   	for _, col2 := range table2.Columns {
   		if fuzzyMatch(col1.Name, col2.Name) {
   			if !strings.EqualFold(col1.DataType, col2.DataType) {
   				mismatches["Data Types"] = fmt.Sprintf("Mismatched: %s -> %s != %s", col1.Name, col1.DataType, col2.DataType)
   			}
   		}
   	}
   }

   return mismatches
}

func main() {
   // Sample data with evolved schemas
   table1 := Table{
   	Name: "Orders",
   	Columns: []Column{
   		{Name: "order_id", DataType: "INT"},
   		{Name: "cust