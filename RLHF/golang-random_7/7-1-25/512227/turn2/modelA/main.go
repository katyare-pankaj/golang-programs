package main

import (
	"fmt"
<<<<<<< HEAD
	"reflect"
)

type Employee struct {
	Name       string
	Age        int
	Department string
	Salary     float32
}

func compareEmployeeSlices(slice1, slice2 []Employee) {
	// Create maps for quick lookup by Name
	map1 := make(map[string]Employee)
	map2 := make(map[string]Employee)

	for _, emp := range slice1 {
		map1[emp.Name] = emp
	}

	for _, emp := range slice2 {
		map2[emp.Name] = emp
	}

	// Find unique employees in slice1
	uniqueToSlice1 := []string{}
	for name := range map1 {
		if _, ok := map2[name]; !ok {
			uniqueToSlice1 = append(uniqueToSlice1, name)
		}
	}

	// Find unique employees in slice2
	uniqueToSlice2 := []string{}
	for name := range map2 {
		if _, ok := map1[name]; !ok {
			uniqueToSlice2 = append(uniqueToSlice2, name)
		}
	}

	// Print unique employees in slice1
	if len(uniqueToSlice1) > 0 {
		fmt.Println("Employees unique to slice1:")
		for _, name := range uniqueToSlice1 {
			fmt.Println(map1[name])
		}
	}

	// Print unique employees in slice2
	if len(uniqueToSlice2) > 0 {
		fmt.Println("\nEmployees unique to slice2:")
		for _, name := range uniqueToSlice2 {
			fmt.Println(map2[name])
		}
	}

	// Find and print differences in field values for matching records by Name
	fmt.Println("\nDifferences in matching employees:")
	for name := range map1 {
		if emp1, ok := map1[name]; ok {
			if emp2, ok := map2[name]; ok {
				// Compare fields using reflection
				typeOfEmp := reflect.TypeOf(emp1)
				valueOfEmp1 := reflect.ValueOf(emp1)
				valueOfEmp2 := reflect.ValueOf(emp2)

				for i := 0; i < typeOfEmp.NumField(); i++ {
					field := typeOfEmp.Field(i)
					if field.Name == "Name" {
						continue // Skip Name field as it's used for matching
					}

					value1 := valueOfEmp1.Field(i)
					value2 := valueOfEmp2.Field(i)

					if value1.Interface() != value2.Interface() {
						fmt.Printf("\nEmployee Name: %s\n", name)
						fmt.Printf("Field: %s, Value in slice1: %v, Value in slice2: %v\n", field.Name, value1.Interface(), value2.Interface())
					}
				}
			}
		}
=======
	"math/rand"
	"sync"
	"time"
)

// LogMessage struct holds log message data
type LogMessage struct {
	Level   string
	Message string
}

// GenerateLogs creates a goroutine to generate log messages
func GenerateLogs(msgChan chan LogMessage, wg *sync.WaitGroup) {
	defer wg.Done()

	logLevels := []string{"INFO", "WARNING", "ERROR", "DEBUG"}
	for {
		// Generate random log level and message
		level := logLevels[rand.Intn(len(logLevels))]
		msg := fmt.Sprintf("Log %d: %s - %s", rand.Intn(1000), level, fmt.Sprintf("Random message %d", rand.Intn(1000)))
		// Send the log message to the channel
		msgChan <- LogMessage{Level: level, Message: msg}
		// Simulate a random interval between log messages
		time.Sleep(time.Duration(rand.Intn(1000)+50) * time.Millisecond)
	}
}

// ProcessLogs is the worker goroutine that processes incoming log messages
func ProcessLogs(msgChan chan LogMessage) {
	for msg := range msgChan {
		fmt.Printf("Processed log: %s: %s\n", msg.Level, msg.Message)
>>>>>>> c5a2dceff1671c8fb06ad48027433ec567c4d577
	}
}

func main() {
<<<<<<< HEAD
	slice1 := []Employee{
		{Name: "Alice", Age: 30, Department: "Engineering", Salary: 70000.0},
		{Name: "Bob", Age: 25, Department: "Design", Salary: 50000.0},
		{Name: "Charlie", Age: 35, Department: "Marketing", Salary: 60000.0},
	}

	slice2 := []Employee{
		{Name: "Alice", Age: 31, Department: "Engineering", Salary: 72000.0},
		{Name: "David", Age: 28, Department: "Sales", Salary: 55000.0},
		{Name: "Charlie", Age: 35, Department: "Marketing", Salary: 62000.0},
	}

	compareEmployeeSlices(slice1, slice2)
=======
	// Initialize random number generator
	rand.Seed(time.Now().UnixNano())

	// Create a channel to pass log messages
	msgChan := make(chan LogMessage)

	// WaitGroup to ensure all generator goroutines are completed
	var wg sync.WaitGroup
	const numGenerators = 5 // Number of generator goroutines

	// Start generator goroutines
	for i := 0; i < numGenerators; i++ {
		wg.Add(1)
		go GenerateLogs(msgChan, &wg)
	}

	// Start the worker goroutine
	go ProcessLogs(msgChan)

	// Wait for all generators to complete
	wg.Wait()

	// Gracefully terminate the worker by closing the channel
	close(msgChan)
>>>>>>> c5a2dceff1671c8fb06ad48027433ec567c4d577
}
