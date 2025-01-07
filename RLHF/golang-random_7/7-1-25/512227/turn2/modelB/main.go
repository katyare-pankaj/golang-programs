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
	Salary     float64
}

func main() {
	emp1 := []Employee{
		{Name: "Alice", Age: 30, Department: "Engineering", Salary: 75000.0},
		{Name: "Bob", Age: 35, Department: "Sales", Salary: 60000.0},
		{Name: "Charlie", Age: 40, Department: "Operations", Salary: 50000.0},
	}

	emp2 := []Employee{
		{Name: "Alice", Age: 30, Department: "Engineering", Salary: 75000.0},
		{Name: "Bob", Age: 36, Department: "Sales", Salary: 62000.0},
		{Name: "David", Age: 32, Department: "Marketing", Salary: 68000.0},
	}

	compareEmployees(emp1, emp2)
}

func compareEmployees(emp1, emp2 []Employee) {
	fmt.Println("Comparing Employees:")
	fmt.Println("--------------------")

	// Create a map to store employees from emp1 for efficient lookups
	emp1Map := make(map[string]Employee)
	for _, emp := range emp1 {
		emp1Map[emp.Name] = emp
	}

	for _, emp2 := range emp2 {
		if emp1, ok := emp1Map[emp2.Name]; ok {
			// Found matching employee in emp1, compare fields
			compareStructs(emp1, emp2)
		} else {
			// Employee not found in emp1, print as a difference
			fmt.Printf("Employee %s not found in emp1.\n", emp2.Name)
		}
	}
}

func compareStructs(emp1, emp2 Employee) {
	v1 := reflect.ValueOf(emp1)
	v2 := reflect.ValueOf(emp2)
	t := v1.Type()

	for i := 0; i < t.NumField(); i++ {
		fieldName := t.Field(i).Name
		field1 := v1.Field(i).Interface()
		field2 := v2.Field(i).Interface()

		if field1 != field2 {
			fmt.Printf("Difference in %s: %v vs %v\n", fieldName, field1, field2)
		}
	}
=======
	"math/rand"
	"sync"
	"time"
)

const (
	numGenerators = 5
	logBufferSize = 100
	logInterval   = 100 * time.Millisecond
)

func generateLogs(wg *sync.WaitGroup, logCh chan string) {
	defer wg.Done()
	for {
		log := fmt.Sprintf("Log message generated at %v", time.Now())
		logCh <- log
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	}
}
func processLogs(wg *sync.WaitGroup, logCh chan string) {
	defer wg.Done()
	for log := range logCh {
		fmt.Println(log)
		// Simulate processing time
		time.Sleep(logInterval)
	}
}
func main() {
	var wg sync.WaitGroup
	logCh := make(chan string, logBufferSize)
	wg.Add(numGenerators + 1)
	// Start generating log messages
	for i := 0; i < numGenerators; i++ {
		go generateLogs(&wg, logCh)
	}
	// Start processing log messages
	go processLogs(&wg, logCh)
	wg.Wait()
>>>>>>> c5a2dceff1671c8fb06ad48027433ec567c4d577
}
