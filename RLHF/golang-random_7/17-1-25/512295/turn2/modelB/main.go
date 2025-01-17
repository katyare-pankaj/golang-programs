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
		params[key] = values[0]
	}
	return params, nil
}

// CustomParser is another concrete strategy with custom parsing logic
type CustomParser struct{}

// Parse implements the QueryParserStrategy interface for CustomParser
func (cp *CustomParser) Parse(query string) (map[string]string, error) {
	// Your custom parsing logic goes here
	params := make(map[string]string)
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

// ParseStrategyFactory defines a factory method to create a QueryParserStrategy based on a name
type ParseStrategyFactory interface {
	CreateStrategy(name string) QueryParserStrategy
}

// DefaultParseStrategyFactory implements ParseStrategyFactory to create the default strategy
type DefaultParseStrategyFactory struct{}

// CreateStrategy implements the ParseStrategyFactory interface for DefaultParseStrategyFactory
func (f *DefaultParseStrategyFactory) CreateStrategy(name string) QueryParserStrategy {
	switch name {
	case "simple":
		return &SimpleParser{}
	case "urlvalues":
		return &URLValuesParser{}
	case "custom":
		return &CustomParser{}
	default:
		// Return a default strategy or handle error
		return &SimpleParser{}
	}
}

// main function to demonstrate the Strategy pattern
func main() {
	query := "name=JohnDoe&age=30&country=USA"

	// Create a factory to create parsing strategies
	factory := &DefaultParseStrategyFactory{}

	context := &QueryParserContext{}

	// Use different strategies at runtime based on user input or configuration
	strategyNames := []string{"simple", "urlvalues", "custom", "invalid"}
	for _, strategyName := range strategyNames {
		fmt.Println("\nUsing", strategyName, "strategy:")
		strategy := factory.CreateStrategy(strategyName)
		context.SetStrategy(strategy)
		params, err := context.Parse(query)
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Println("Parsed Parameters:", params)
		}
	}
}
