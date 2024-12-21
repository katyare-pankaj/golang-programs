package main

import "fmt"

type DataProcessor func(interface{}) error

func ProcessDataWithCallback(data []interface{}, callback DataProcessor) error {
	for _, item := range data {
		if err := callback(item); err != nil {
			return err
		}
	}
	return nil
}

func logData(data interface{}) error {
	switch val := data.(type) {
	case int:
		fmt.Println("Logging integer:", val)
	case string:
		fmt.Println("Logging string:", val)
	default:
		return fmt.Errorf("unexpected data type: %T", data)
	}
	return nil
}

func main() {
	data := []interface{}{1, 2, "three", 4.5, true}
	err := ProcessDataWithCallback(data, logData)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
