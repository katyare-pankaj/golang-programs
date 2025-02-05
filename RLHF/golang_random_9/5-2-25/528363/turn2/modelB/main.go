// main.go
package main

import (
    "fmt"
    "io/ioutil"
    "log"
    "os"
    "plugin"
    "strings"
)

func main() {
    if len(os.Args) != 3 {
        fmt.Println("Usage: go run main.go <plugin-name> <file-path>")
        return
    }

    pluginName := os.Args[1]
    filePath := os.Args[2]

    p, err := loadPlugin(pluginName)
    if err != nil {
        log.Fatalf("Error loading plugin: %v", err)
    }

    if err := processFile(p, filePath); err != nil {
        log.Fatalf("Error processing file: %v", err)
    }

    fmt.Println("File processed successfully")
}

func loadPlugin(pluginName string) (TextProcessor, error) {
    // Add the correct plugin extension based on the OS
    pluginPath := pluginName + string(os.PathSeparator) + pluginName + ".so"

    p, err := plugin.Open(pluginPath)
    if err != nil {
        return nil, err
    }

    symTextProcessor, err := p.Lookup("TextProcessor")
    if err != nil {
        return nil, err
    }

    var textProc TextProcessor
    textProc, ok := symTextProcessor.(TextProcessor)
    if !ok {
        return nil, fmt.Errorf("unexpected type from module symbol")
    }

    return textProc, nil
}

func processFile(tp TextProcessor, filePath string) error {
    data, err := ioutil.ReadFile(filePath)
    if err != nil {
        return err
    }