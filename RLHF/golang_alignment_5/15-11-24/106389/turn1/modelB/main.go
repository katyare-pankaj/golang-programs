package main

import (
	"fmt"
	"sync"
	"time"
)

type Task interface {
	Execute()
}

type PrimitiveTask struct {
	name string
	work func()
}

func (t *PrimitiveTask) Execute() {
	fmt.Println("Executing primitive task:", t.name)
	t.work()
}

func newPrimitiveTask(name string, work func()) *PrimitiveTask {
	return &PrimitiveTask{name: name, work: work}
}

type CompositeTask struct {
	tasks []Task
}

func (t *CompositeTask) Execute() {
	fmt.Println("Executing composite task:")
	for _, task := range t.tasks {
		task.Execute()
	}
}

func (t *CompositeTask) AddTask(task Task) {
	t.tasks = append(t.tasks, task)
}

func parallelExecute(tasks []Task) {
	var wg sync.WaitGroup
	wg.Add(len(tasks))

	for _, task := range tasks {
		go func(t Task) {
			t.Execute()
			wg.Done()
		}(task)
	}
	wg.Wait()
}

func main() {
	// Simple sequential task
	task1 := newPrimitiveTask("Work 1", func() {
		time.Sleep(2 * time.Second)
	})

	// Another simple sequential task
	task2 := newPrimitiveTask("Work 2", func() {
		time.Sleep(1 * time.Second)
	})

	// Task composite to run Work1 and Work2 sequentially
	sequentialTasks := &CompositeTask{tasks: []Task{task1, task2}}

	// Creating a composite of parallel tasks
	task3 := newPrimitiveTask("Work 3", func() {
		time.Sleep(3 * time.Second)
	})

	task4 := newPrimitiveTask("Work 4", func() {
		time.Sleep(2 * time.Second)
	})

	parallelTasks := &CompositeTask{tasks: []Task{task3, task4}}

	// Running parallel composite task
	fmt.Println("Running parallel tasks:")
	startTime := time.Now()
	parallelExecute(parallelTasks.tasks)
	fmt.Println("Parallel tasks finished in:", time.Since(startTime))

	// Running sequential composite task
	fmt.Println("\nRunning sequential tasks:")
	sequentialTasks.Execute()
}
