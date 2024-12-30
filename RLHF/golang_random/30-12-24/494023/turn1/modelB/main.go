package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"sync"
)

func countWords(filename string, wg *sync.WaitGroup, wordCounts *sync.Map) {
	defer wg.Done()
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Println("Error reading file:", err)
		return
	}
	words := strings.Fields(string(data))
	for _, word := range words {
		value, _ := wordCounts.LoadOrStore(word, 0)
		wordCounts.Store(word, value.(int)+1)
	}
}

func main() {
	filenames := []string{"file1.txt", "file2.txt", "file3.txt"}
	var wg sync.WaitGroup
	wordCounts := new(sync.Map)
	wg.Add(len(filenames))
	for _, filename := range filenames {
		go countWords(filename, &wg, wordCounts)
	}
	wg.Wait()
	var totalWords int
	wordCounts.Range(func(key, value interface{}) bool {
		totalWords += value.(int)
		return true
	})
	fmt.Println("Total words:", totalWords)
}
