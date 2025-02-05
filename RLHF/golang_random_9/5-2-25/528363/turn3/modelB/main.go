// main.go
package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"plugin"
	"strings"
)

func LoadPlugins(pluginPaths []string) ([]textprocessor.TextProcessor, error) {
	var processors []textprocessor.TextProcessor

	for _, pluginPath := range pluginPaths {
		p, err := plugin.Open(pluginPath)
		if err != nil {
			return nil, fmt.Errorf("error opening plugin '%s': %w", pluginPath, err)
		}

		symTextProcessor, err := p.Lookup("TextProcessorPlugin")
		if err != nil {
			return nil, fmt.Errorf("error looking up TextProcessorPlugin in '%s': %w", pluginPath, err)
		}

		var processor textprocessor.TextProcessor
		processor, ok := symTextProcessor.(textprocessor.TextProcessor)
		if !ok {
			return nil, errors.New("unexpected type from module symbol")
		}
		processors = append(processors, processor)
	}

	return processors, nil