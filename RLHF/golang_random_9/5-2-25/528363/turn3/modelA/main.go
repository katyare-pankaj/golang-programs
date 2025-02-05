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

// LoadPlugin loads a plugin and returns a TextProcessor instance
func LoadPlugin(pluginPath string) (textprocessor.TextProcessor, error) {
	p, err := plugin.Open(pluginPath)
	if err != nil {
		return nil, fmt.Errorf("error opening plugin: %w", err)
	}

	symTextProcessor, err := p.Lookup("TextProcessorPlugin")
	if err != nil {
		return nil, fmt.Errorf("error looking up TextProcessorPlugin: %w", err)
	}

	processor, ok := symTextProcessor.(textprocessor.TextProcessor)
	if !ok {
		return nil, fmt.Errorf("unexpected type from module symbol")
	}

	return processor, nil
}

func main() {
	if len(os.Args) < 4 {
		log.Fatalf("Usage: %s <plugin-path1> <plugin-path2> ... <input-file> <output-file>\n", os.Args[0])
	}

	pluginPaths := os.Args[1 : len(os.Args)-2]
	inputFile := os.Args[len(os.Args)-2]
	outputFile := os.Args[len(os.Args)-1]

	// Load the input file
	inputData, err := ioutil.ReadFile(inputFile)
	if err != nil {
		log.Fatalf("Error reading input file: %v", err)
	}

	// Process the text with each plugin sequentially
	processedText := string(inputData)
	for _, pluginPath := range pluginPaths {
		processor, err := LoadPlugin(pluginPath)
		if err != nil {
			log.Fatalf("Failed to load plugin: %v", err)
		}

		processedText, err = processor.Process(processedText)
		if err != nil {
			log.Fatalf("Error processing text with plugin %s: %v", pluginPath, err)
		}
	}

	// Write the processed text to the output file
	err = ioutil.WriteFile(outputFile, []byte(processedText), 0644)
	if err != nil {
		log.Fatalf("Error writing to output file: %v", err)
	}

	fmt.Println("Text processing completed successfully.")
}
