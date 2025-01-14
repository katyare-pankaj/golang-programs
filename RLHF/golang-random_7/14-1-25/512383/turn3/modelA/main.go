package main

import (
	"fmt"
	"sync"
	"time"
)

type Cell struct {
	value   int
	readers int
	mu      sync.Mutex
	wg      sync.WaitGroup
}

type Table struct {
	cells map[int]*Cell
}

func (t *Table) CreateCell(key int) {
	t.cells[key] = &Cell{}
}

func (c *Cell) Read() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.readers++
	c.wg.Add(1)
	defer c.wg.Done()
	return c.value
}

func (c *Cell) Write(value int) {
	c.mu.Lock()
	defer c.mu.Unlock()
	// Wait for all readers to finish
	c.wg.Wait()
	c.value = value
}

func (t *Table) GetCell(key int) *Cell {
	return t.cells[key]
}

func (t *Table) UpdateCell(key int, value int) {
	cell := t.GetCell(key)
	if cell != nil {
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
