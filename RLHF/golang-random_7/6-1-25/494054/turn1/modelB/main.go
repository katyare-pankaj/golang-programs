package main

import (
	"fmt"
	"strings"
)

type TransformRule struct {
	Name     string
	Template string
	Args     []interface{}
}

func ApplyTransformPipeline(input string, rules []TransformRule) string {
	result := input
	for _, rule := range rules {
		result = fmt.Sprintf(rule.Template, append([]interface{}{result}, rule.Args...)...)
	}
	return result
}

func main() {
	input := "Hello, World!"
	rules := []TransformRule{
		{
			Name:     "Uppercase",
			Template: "%s",
			Args:     []interface{}{strings.ToUpper},
		},
		{
			Name:     "Reverse",
			Template: "%s",
			Args: []interface{}{func(s string) string {
				var reversed []rune
				for _, r := range s {
					reversed = append([]rune{r}, reversed...)
				}
				return string(reversed)
			}},
		},
		{
			Name:     "Add Prefix",
			Template: "%s",
			Args:     []interface{}{"Prefix: "},
		},
	}

	transformed := ApplyTransformPipeline(input, rules)
	fmt.Println("Transformed String:", transformed)
}
