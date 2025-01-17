package main

import (
	"fmt"
	"net/url"
)

// ParserStrategy defines the interface for parsing URL query parameters
type ParserStrategy interface {
	Parse(values url.Values) (map[string]string, error)
}

// SimpleParser is a basic strategy that parses query parameters as-is
type SimpleParser struct{}

func (s *SimpleParser) Parse(values url.Values) (map[string]string, error) {
	result := make(map[string]string)
	for key, val := range values {
		result[key] = val[0]
	}
	return result, nil
}

// CustomParser is a more complex strategy that performs custom parsing
type CustomParser struct {
	dateField string // The field to be parsed as a date
}

func (s *CustomParser) Parse(values url.Values) (map[string]string, error) {
	result := make(map[string]string)
	for key, val := range values {
		switch key {
		case s.dateField:
			// Custom parsing logic for date field, e.g., using time package
			// For simplicity, let's just append a suffix
			result[key] = val[0] + "-custom"
		default:
			result[key] = val[0]
		}
	}
	return result, nil
}

// ParserContext uses the ParserStrategy to parse query parameters
type ParserContext struct {
	strategy ParserStrategy
}

func NewParserContext(strategy ParserStrategy) *ParserContext {
	return &ParserContext{strategy: strategy}
}

func (pc *ParserContext) SetStrategy(strategy ParserStrategy) {
	pc.strategy = strategy
}

func (pc *ParserContext) ParseQuery(query string) (map[string]string, error) {
	u, err := url.Parse(query)
	if err != nil {
		return nil, err
	}
	return pc.strategy.Parse(u.Query())
}

func main() {
	simpleParser := NewParserContext(&SimpleParser{})
	customParser := NewParserContext(&CustomParser{dateField: "date"})

	// Example URL query strings
	query1 := "?name=Alice&age=25"
	query2 := "?name=Bob&age=30&date=2023-07-28"

	// Use simple parser
	params1, err := simpleParser.ParseQuery(query1)
	if err != nil {
		fmt.Println("Error parsing query 1:", err)
		return
	}
	fmt.Println("Query 1 parsed by SimpleParser:")
	for key, val := range params1 {
		fmt.Printf("%s: %s\n", key, val)
	}
	fmt.Println()

	// Use custom parser
	params2, err := customParser.ParseQuery(query2)
	if err != nil {
		fmt.Println("Error parsing query 2:", err)
		return
	}
	fmt.Println("Query 2 parsed by CustomParser:")
	for key, val := range params2 {
		fmt.Printf("%s: %s\n", key, val)
	}
}
