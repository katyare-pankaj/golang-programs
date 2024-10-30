package main

import (
	"fmt"
	"log"

	"github.com/garyburd/redigo/redis"
)

func main() {
	// Connect to Redis in-memory database
	c, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		log.Fatalf("Error connecting to Redis: %v", err)
	}
	defer c.Close()

	// Load some sample data into Redis (Key: PropertyID, Value: LastSoldPrice)
	err = loadSampleData(c)
	if err != nil {
		log.Fatalf("Error loading sample data: %v", err)
	}

	// Example 1: Calculate the average price of properties based on last sold data
	avgPrice, err := calculateAveragePrice(c)
	if err != nil {
		log.Fatalf("Error calculating average price: %v", err)
	}
	fmt.Println("Average Last Sold Price:", avgPrice)

	// Example 2: Find the 5 most expensive properties in the area
	top5ExpensiveProperties, err := getTop5ExpensiveProperties(c)
	if err != nil {
		log.Fatalf("Error getting top 5 expensive properties: %v", err)
	}
	fmt.Println("Top 5 Most Expensive Properties:")
	for _, property := range top5ExpensiveProperties {
		fmt.Println(property)
	}
}

func loadSampleData(c redis.Conn) error {
	data := map[string]int{
		"PROP1": 250000,
		"PROP2": 300000,
		"PROP3": 280000,
		"PROP4": 400000,
		"PROP5": 320000,
		// Add more sample data as needed
	}

	for key, value := range data {
		_, err := c.Do("SET", key, value)
		if err != nil {
			return err
		}
	}
	return nil
}

func calculateAveragePrice(c redis.Conn) (float64, error) {
	values, err := redis.Ints(c.Do("VALS", "*"))
	if err != nil {
		return 0, err
	}

	totalPrice := 0
	for _, price := range values {
		totalPrice += price
	}

	return float64(totalPrice) / float64(len(values)), nil
}

func getTop5ExpensiveProperties(c redis.Conn) ([]string, error) {
	values, err := redis.Ints(c.Do("VALS", "*"))
	if err != nil {
		return nil, err
	}

	properties := make([]string, len(values))
