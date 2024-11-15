package main

import (
	"fmt"
)

type Task interface {
	Add(Task)
	Remove(Task)
	IsLeaf() bool
	Execute()
	String() string
}

type leafTask struct {
	name string
}

func (t *leafTask) Add(Task) {}

func (t *leafTask) Remove(Task) {}

func (t *leafTask) IsLeaf() bool {
	return true
}

func (t *leafTask) Execute() {
	fmt.Println("Executing leaf task:", t.name)
}

func (t *leafTask) String() string {
	return t.name
}

type compositeTask struct {
	subtasks []Task
}

func (t *compositeTask) Add(task Task) {
	t.subtasks = append(t.subtasks, task)
}

func (t *compositeTask) Remove(task Task) {
	for i, subtask := range t.subtasks {
		if subtask == task {
			t.subtasks = append(t.subtasks[:i], t.subtasks[i+1:]...)
			return
		}
	}
}

func (t *compositeTask) IsLeaf() bool {
	return false
}

func (t *compositeTask) Execute() {
	fmt.Println("Executing composite task:")
	for _, subtask := range t.subtasks {
		subtask.Execute()
	}
}

func (t *compositeTask) String() string {
	return "Composite(" + fmt.Sprint(t.subtasks) + ")"
}

func main() {
	root := new(compositeTask)

	task1 := &leafTask{name: "Task 1"}
	task2 := &leafTask{name: "Task 2"}
	task3 := &leafTask{name: "Task 3"}
	task4 := &leafTask{name: "Task 4"}

	subtask1 := new(compositeTask)
	subtask1.Add(task1)
	subtask1.Add(task2)

	root.Add(subtask1)
	root.Add(task3)
	root.Add(task4)

	fmt.Println("Task hierarchy:", root)
	root.Execute()
}
