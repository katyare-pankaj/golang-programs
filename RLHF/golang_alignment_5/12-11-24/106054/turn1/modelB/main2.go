package main

import (
	"container/heap"
	"fmt"
)

// Define a Task with priority
type Task struct {
	Value    interface{}
	Priority int
}

// Implement heap.Interface
type PriorityQueue []*Task

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].Priority > pq[j].Priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*Task)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil // avoid memory leak
	*pq = old[0 : n-1]
	return item
}

func main() {
	// Create a priority queue
	pq := make(PriorityQueue, 0)
	heap.Init(&pq)

	// Example tasks with priorities
	heap.Push(&pq, &Task{Value: "High Task", Priority: 10})
	heap.Push(&pq, &Task{Value: "Medium Task", Priority: 5})
	heap.Push(&pq, &Task{Value: "Low Task", Priority: 1})

	// Process tasks based on priority
	for pq.Len() > 0 {
		task := heap.Pop(&pq).(*Task)
		fmt.Println("Processing task:", task.Value, "with priority:", task.Priority)
	}
}
