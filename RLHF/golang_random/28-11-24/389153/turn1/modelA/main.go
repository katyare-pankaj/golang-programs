package main

import (
	"fmt"
	"time"
)

type DataSource interface {
	Next() (string, error)
}
type Processor interface {
	Process(data string) string
}
type InMemoryDataSource struct {
	data  []string
	index int
}

func NewInMemoryDataSource(data []string) *InMemoryDataSource {
	return &InMemoryDataSource{data: data, index: 0}
}

func (ds *InMemoryDataSource) Next() (string, error) {
	if ds.index >= len(ds.data) {
		return "", fmt.Errorf("no more data")
	}
	data := ds.data[ds.index]
	ds.index++
	return data, nil
}

type SimpleProcessor struct{}

func NewSimpleProcessor() *SimpleProcessor {
	return &SimpleProcessor{}
}

func (p *SimpleProcessor) Process(data string) string {
	return fmt.Sprintf("Processed: %s", data)
}
func AnalyticsPipeline(dataSource DataSource, processor Processor) error {
	for {
		data, err := dataSource.Next()
		if err != nil {
			return err
		}
		processedData := processor.Process(data)
		fmt.Println(processedData)
		time.Sleep(1 * time.Second) // Simulate delay between data processing
	}
}

func main() {
	data := []string{"DataPoint1", "DataPoint2", "DataPoint3"}
	dataSource := NewInMemoryDataSource(data)
	processor := NewSimpleProcessor()

	err := AnalyticsPipeline(dataSource, processor)
	if err != nil {
		fmt.Println("Error in analytics pipeline:", err)
	}
}
