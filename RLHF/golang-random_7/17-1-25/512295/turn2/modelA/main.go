package main

import (
	"errors"
	"fmt"
	"net/url"
	"strings"
)

// QueryParserStrategy is an interface for different parsing strategies
type QueryParserStrategy interface {
	Parse(query string) (map[string]string, error)
}

// SimpleParser implements a basic strategy for parsing query strings
type SimpleParser struct{}

func (sp *SimpleParser) Parse(query string) (map[string]string, error) {
	params := make(map[string]string)
	if query == "" {
		return params, nil
	}
	pairs := strings.Split(query, "&")
	for _, pair := range pairs {
		kv := strings.Split(pair, "=")
		if len(kv) != 2 {
			return nil, errors.New("invalid query parameter pair: " + pair)
		}
		params[kv[0]] = kv[1]
	}
	return params, nil
}

// URLValuesParser uses the net/url package to parse query strings
type URLValuesParser struct{}

func (up *URLValuesParser) Parse(query string) (map[string]string, error) {
	params := make(map[string]string)
	parsedValues, err := url.ParseQuery(query)
	if err != nil {
		return nil, err
	}
	for key, values := range parsedValues {
		if len(values) > 0 {
			params[key] = values[0]
		}
	}
	return params, nil
}

// FlexibleParserContext is the context for parsing using different strategies
type FlexibleParserContext struct {
	strategy QueryParserStrategy
}

// SetStrategy changes the current parsing strategy
func (fpc *FlexibleParserContext) SetStrategy(strategy QueryParserStrategy) {
	fpc.strategy = strategy
}

// Parse uses the current strategy to parse the given query
func (fpc *FlexibleParserContext) Parse(query string) (map[string]string, error) {
	return fpc.strategy.Parse(query)
}

func main() {
	query := "name=JaneDoe&age=25&city=NewYork"

	context := &FlexibleParserContext{}

	// Using SimpleParser strategy
	context.SetStrategy(&SimpleParser{})
	params, err := context.Parse(query)
	if err != nil {
		fmt.Println("Error with SimpleParser:", err)
	} else {
		fmt.Println("SimpleParser:", params)
	}

	// Switching to URLValuesParser strategy
	context.SetStrategy(&URLValuesParser{})
	params, err = context.Parse(query)
	if err != nil {
		fmt.Println("Error with URLValuesParser:", err)
	} else {
		fmt.Println("URLValuesParser:", params)
	}

	// Example to demonstrate switching strategies based on some condition
	// Here, we are assuming an arbitrary runtime condition
	useAdvancedParsing := true
	if useAdvancedParsing {
		context.SetStrategy(&URLValuesParser{})
	} else {
		context.SetStrategy(&SimpleParser{})
	}
	params, err = context.Parse(query)
	if err != nil {
		fmt.Println("Error with chosen parser:", err)
	} else {
		fmt.Println("Chosen parser:", params)
	}
}
