package main

import (
	"fmt"
	"time"
)

// Task represents a transaction or task for analysis
type Task struct {
	name      string
	startTime time.Time
	endTime   time.Time
	cost      float64
}

// Analyzer provides methods to analyze transaction data
type Analyzer struct {
	tasks []Task
}

// NewAnalyzer creates a new Analyzer instance
func NewAnalyzer() *Analyzer {
	return &Analyzer{tasks: []Task{}}
}

// AddTask adds a new task to the analyzer
func (a *Analyzer) AddTask(task Task) {
	a.tasks = append(a.tasks, task)
}

// CalculateTotalCost calculates the total cost of all tasks
func (a *Analyzer) CalculateTotalCost() float64 {
	totalCost := 0.0
	for _, task := range a.tasks {
		totalCost += task.cost
	}
	return totalCost
}

func main() {
	// Agile principle: Independent development teams (small modules)
	taskAnalyzer := NewAnalyzer()

	// Agile principle: Incremental Development
	addTasks(taskAnalyzer)

	// Perform analysis
	totalCost := taskAnalyzer.CalculateTotalCost()
	fmt.Printf("Total Cost: $%.2f\n", totalCost)

	// Agile principle: Frequent Releases and Feedback
	// Continuous Improvement can be implemented here by adding user input or new requirements
}

func addTasks(analyzer *Analyzer) {
	// Simulate adding tasks in increments
	analyzer.AddTask(Task{name: "Task 1", startTime: time.Now(), endTime: time.Now().Add(time.Minute), cost: 10.0})
	analyzer.AddTask(Task{name: "Task 2", startTime: time.Now().Add(time.Minute), endTime: time.Now().Add(2 * time.Minute), cost: 20.0})
	// Add more tasks as needed
}
