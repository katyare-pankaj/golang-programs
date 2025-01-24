package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type AppError struct {
	Original error
	Message  string
	Code     int
}

func (e *AppError) Error() string {
	return fmt.Sprintf("Code: %d, Message: %s, Original: %s", e.Code, e.Message, e.Original.Error())
}
func openFile(filename string) (*os.File, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, &AppError{
			Original: err,
			Message:  fmt.Sprintf("Error opening file: %s", filename),
			Code:     404,
		}
	}
	return f, nil
}
func readFile(file *os.File) (string, error) {
	b, err := ioutil.ReadAll(file)
	if err != nil {
		return "", &AppError{
			Original: err,
			Message:  "Error reading file",
			Code:     500,
		}
	}
	return string(b), nil
}
func processData(data string) (string, error) {
	if data == "" {
		return "", &AppError{
			Message: "Data is empty",
			Code:    400,
		}
	}
	result := strings.ToUpper(data)
	return result, nil
}
func main() {
	filename := "example.txt"
	f, err := openFile(filename)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer f.Close()
	data, err := readFile(f)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	result, err := processData(data)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Result:", result)
}
