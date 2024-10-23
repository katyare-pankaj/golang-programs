package main

import (
	"fmt"
)

type Schema struct {
	Fields []Field
}

type Field struct {
	Name string
	Type string
}

func schemaSimilarity(schema1, schema2 Schema) float64 {
	commonFields := 0
	totalFields := len(schema1.Fields) + len(schema2.Fields)

	for _, field1 := range schema1.Fields {
		for _, field2 := range schema2.Fields {
			if field1.Name == field2.Name && field1.Type == field2.Type {
				commonFields++
				break
			}
		}
	}

	return float64(commonFields) / float64(totalFields)
}

func main() {
	schema1 := Schema{
		Fields: []Field{
			{Name: "id", Type: "int"},
			{Name: "name", Type: "string"},
			{Name: "age", Type: "int"},
		},
	}

	schema2 := Schema{
		Fields: []Field{
			{Name: "ID", Type: "integer"},
			{Name: "full_name", Type: "string"},
			{Name: "age", Type: "int"},
			{Name: "location", Type: "string"},
		},
	}

	similarity := schemaSimilarity(schema1, schema2)
	fmt.Println("Schema Similarity:", similarity) // Output: Schema Similarity: 0.6
}
