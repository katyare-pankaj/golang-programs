// main.go
package main

import (
	"fmt"
	"go-programs/RLHF/golang_random_9/5-2-25/528363/turn2/modelA/textprocessor"
	"io/ioutil"
	"log"
	"os"
	"plugin"
)

// LoadPlugin dynamically loads a plugin by its path
func LoadPlugin(pluginPath string) (textprocessor.TextProcessor, error) {
	p, err := plugin.Open(pluginPath)
	if err != nil {
		return nil, fmt.Errorf("error opening plugin: %w", err)
	}

	symTextProcessor, err := p.Lookup("TextProcessorPlugin")
	if err != nil {
		return nil, fmt.Errorf("error looking up TextProcessorPlugin: %w", err)
	}

	var processor textprocessor.TextProcessor
	processor, ok := symTextProcessor.(textprocessor.TextProcessor)
	if !ok {
		return nil, fmt.Errorf("unexpected type from module symbol")
	}

	return processor, nil
}

func main() {
	if len(os.Args) != 4 {
		log.Fatalf("Usage: %s <plugin-path> <input-file> <output-file>\n", os.Args[0])
	}

	pluginPath := os.Args[1]
	inputFile := os.Args[2]
	outputFile := os.Args[3]

	processor, err := LoadPlugin(pluginPath)
	if err != nil {
		log.Fatalf("Failed to load plugin: %v", err)
	}

	inputData, err := ioutil.ReadFile(inputFile)
	if err != nil {
		log.Fatalf("Error reading input file: %v", err)
	}

	processedText, err := processor.Process(string(inputData))
	if err != nil {
		log.Fatalf("Error processing text: %v", err)
	}

	err = ioutil.WriteFile(outputFile, []byte(processedText), 0644)
	if err != nil {
		log.Fatalf("Error writing to output file: %v", err)
	}

	fmt.Println("Text processing completed successfully.")
}
