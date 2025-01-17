package main

import (
	"fmt"
	"net/url"
	"strings"
)

type QueryParserStrategy interface {
	Parse(query string) (map[string]string, error)
}

type SimpleParser struct{}

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

type URLValuesParser struct{}

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

type QueryParserContext struct {
	strategy QueryParserStrategy
}

func (qpc *QueryParserContext) SetStrategy(strategy QueryParserStrategy) {
	qpc.strategy = strategy
}

func (qpc *QueryParserContext) Parse(query string) (map[string]string, error) {
	return qpc.strategy.Parse(query)
}

func main() {
	query := "name=JohnDoe&age=30&country=USA"

	context := &QueryParserContext{}

	fmt.Println("Using SimpleParser Strategy:")
	context.SetStrategy(&SimpleParser{})
	params, err := context.Parse(query)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Parsed Parameters:", params)
	}

	fmt.Println("\nUsing URLValuesParser Strategy:")
	context.SetStrategy(&URLValuesParser{})
	params, err = context.Parse(query)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Parsed Parameters:", params)
	}
}
