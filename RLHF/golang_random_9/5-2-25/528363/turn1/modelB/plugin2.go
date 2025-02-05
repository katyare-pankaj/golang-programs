package main

import (
	"fmt"
	"io/ioutil"
)

type reverseProcessor struct{}

func (rp *reverseProcessor) Process(filename string) error {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	reversedData := string(data)[:0]
	for i := len(data) - 1; i >= 0; i-- {
		reversedData += string(data[i])
	}
	err = ioutil.WriteFile(filename, []byte(reversedData), 0644)
	if err != nil {
		return err
	}
	fmt.Println("Reversed file content.")
	return nil
}

var FileProcessor FileProcessor = &reverseProcessor{}
