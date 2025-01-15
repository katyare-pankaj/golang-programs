package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Cell struct {
	value int
	mu    sync.RWMutex
}

type Table struct {
	cells map[int]*Cell
	mu    sync.Mutex
	wg    sync.WaitGroup
}

func NewTable() *Table {
	return &Table{
		cells: make(map[int]*Cell),
	}
}

func (t *Table) CreateCell(key int) {
	t.mu.Lock()
	defer t.mu.Unlock()
	if _, ok := t.cells[key]; !ok {
		t.cells[key] = &Cell{}
	}
}

func (t *Table) GetCell(key int) *Cell {
	t.mu.RLock()
	defer t.mu.RUnlock()
	return t.cells[key]
}

func (c *Cell) Read() int {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.value
}

func (c *Cell) Write(value int) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value = value
}

func (t *Table) UpdateCell(key int, value int) {
	t.wg.Add(1)
	defer t.wg.Done()
	cell := t.GetCell(key)
	if cell != nil {
		cell.Write(value)
	}
}

func (t *Table) ReadCell(key int) int {
	t.wg.Add(1)
	defer t.wg.Done()
	cell := t.GetCell(key)
	if cell != nil {
		return cell.Read()
	}
	return 0 // Or handle missing cell as appropriate
}

func (t *Table) Wait() {
	t.wg.Wait()
}

func main() {
	table := NewTable()

	// Create some cells concurrently
	for i := 1; i <= 5; i++ {
		go table.CreateCell(i)
	}
	// Allow some time for cell creation to complete
	time.Sleep(time.Second)

	// Reader goroutines
	for i := 1; i <= 3; i++ {
		go func(readerID int) {
			for {
				key := rand.Intn(5) + 1 // Select a random cell to read
				value := table.ReadCell(key)
				fmt.Printf("Reader %d reads cell %d: %d\n", readerID, key, value)
				time.Sleep(time.Duration(rand.Intn(3)+1) * time.Second)
			}
		}(i)
	}

	// Writer goroutines
	for i := 1; i <= 3; i++ {
		go func(writerID int) {
			for {
				key := rand.Intn(5) + 1 // Select a random cell to write
				value := table.ReadCell(key)
				table.UpdateCell(key, value+1)
				fmt.Printf("Writer %d updates cell %d: %d\n", writerID, key, value+1)
				time.Sleep(time.Duration(rand.Intn(3)+1) * time.Second)
			}
		}(i)
	}

	// Main goroutine waits for all goroutines to finish
	table.Wait()
}
