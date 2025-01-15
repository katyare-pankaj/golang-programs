package main

import (
	"fmt"
	"sync"
	"time"
)

type Cell struct {
	value int
}

type Table struct {
	cells map[int]*Cell
	rw    sync.RWMutex
}

func (t *Table) CreateCell(key int) {
	t.rw.Lock()
	defer t.rw.Unlock()
	if _, ok := t.cells[key]; !ok {
		t.cells[key] = &Cell{}
	}
}

func (t *Table) GetCell(key int) *Cell {
	t.rw.RLock()
	defer t.rw.RUnlock()
	return t.cells[key]
}

func (c *Cell) Read() int {
	return c.value
}

func (c *Cell) Write(value int) {
	c.value = value
}

func (t *Table) UpdateCell(key int, value int) {
	t.rw.Lock()
	defer t.rw.Unlock()
	if cell := t.GetCell(key); cell != nil {
		cell.Write(value)
	}
}

func main() {
	table := &Table{
		cells: make(map[int]*Cell),
	}

	// Create some cells concurrently
	for i := 1; i <= 5; i++ {
		go func(id int) {
			table.CreateCell(id)
		}(i)
	}

	// Reader goroutines
	for i := 1; i <= 5; i++ {
		go func(readerID int) {
			for {
				key := readerID%5 + 1 // Select a cell to read based on reader ID
				cell := table.GetCell(key)
				if cell != nil {
					fmt.Printf("Reader %d reads cell %d: %d\n", readerID, key, cell.Read())
				}
				time.Sleep(time.Duration(readerID) * time.Millisecond)
			}
		}(i)
	}

	// Writer goroutines
	for i := 1; i <= 5; i++ {
		go func(writerID int) {
			for {
				key := writerID%5 + 1 // Select a cell to write based on writer ID
				table.UpdateCell(key, key*writerID)
				fmt.Printf("Writer %d updates cell %d: %d\n", writerID, key, key*writerID)
				time.Sleep(time.Duration(writerID) * time.Millisecond)
			}
		}(i)
	}

	// You can optionally use `time.After()` to stop the application
	time.Sleep(5 * time.Second)
}
