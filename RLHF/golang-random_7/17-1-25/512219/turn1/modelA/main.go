package main

import (
	"fmt"
	"net/url"
	"strings"
)

// QueryParserStrategy defines the interface for parsing strategies
type QueryParserStrategy interface {
	Parse(query string) (map[string]string, error)
}

// SimpleParser is a concrete strategy that parses the query using simple splitting
type SimpleParser struct{}

// Parse implements the QueryParserStrategy interface for SimpleParser
func (sp *SimpleParser) Parse(query string) (map[string]string, error) {
	params := make(map[string]string)
	pairs := strings.Split(query, "&")
	for _, pair := range pairs {
		kv := strings.Split(pair, "=")
		if len(kv) != 2 {
			return nil, fmt.Errorf("invalid query parameter: %s", pair)
		}
		params[kv[0]] = kv[1]
	}
	return params, nil
}

// URLValuesParser is another concrete strategy using url.ParseQuery
type URLValuesParser struct{}

// Parse implements the QueryParserStrategy interface for URLValuesParser
func (up *URLValuesParser) Parse(query string) (map[string]string, error) {
	params := make(map[string]string)
	parsedValues, err := url.ParseQuery(query)
	if err != nil {
		return nil, err
	}
	for key, values := range parsedValues {
		params[key] = values[0] // only take the first value
	}
	return params, nil
}

// QueryParserContext is the context that uses a QueryParserStrategy
type QueryParserContext struct {
	strategy QueryParserStrategy
}

// SetStrategy allows changing the parsing strategy at runtime
func (qpc *QueryParserContext) SetStrategy(strategy QueryParserStrategy) {
	qpc.strategy = strategy
}

// Parse performs parsing using the current strategy
func (qpc *QueryParserContext) Parse(query string) (map[string]string, error) {
	return qpc.strategy.Parse(query)
}

// main function to demonstrate the Strategy pattern
func main() {
	query := "name=JohnDoe&age=30&country=USA"

	context := &QueryParserContext{}

	// Use SimpleParser strategy
	fmt.Println("Using SimpleParser Strategy:")
	context.SetStrategy(&SimpleParser{})
	params, err := context.Parse(query)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Parsed Parameters:", params)
	}

	// Use URLValuesParser strategy
	fmt.Println("\nUsing URLValuesParser Strategy:")
	context.SetStrategy(&URLValuesParser{})
	params, err = context.Parse(query)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Parsed Parameters:", params)
	}
}
