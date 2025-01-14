package main

import (
	"sync"
)

type Cell struct {
	value     int
	readCount int
	mu        sync.Mutex
	wg        sync.WaitGroup
}

type Table struct {
	cells map[int]*Cell
	mu    sync.Mutex
}

func (t *Table) CreateCell(key int) {
	t.mu.Lock()
	defer t.mu.Unlock()
	if _, ok := t.cells[key]; !ok {
		t.cells[key] = &Cell{}
	}
}

func (t *Table) GetCell(key int) *Cell {
	t.mu.Lock()
	defer t.mu.Unlock()
	return t.cells[key]
}

func (c *Cell) Read() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.readCount++
	defer func() {
		c.readCount--
		if c.readCount == 0 {
			c.wg.Broadcast() // Wake up all writers waiting for this cell
		}
	}()
	return c.value
}

func (c *Cell) Write(value int) {
	c.mu.Lock()
	defer c.mu.Unlock()
	for c.readCount > 0 {
		c.wg.Wait() // Wait for all readers to complete
	}
	c.value = value
}

func (t *Table) UpdateCell(key int, value int) {
	cell := t.GetCell(key)
	if cell != nil {
		cell.Write(value)
	}
}

func (t *Table) ReadCell(key int) int {
	cell := t.GetCell(key)
	if cell != nil {
		return cell.Read()
	}
	return 0
}

func main() {
	table := &Table{
		cells: make(map[int]*Cell),
	}
	// Rest of the code remains the same...
}
