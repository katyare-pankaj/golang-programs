package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type inventory struct {
	sync.Mutex
	items map[string]int
}

func (i *inventory) updateInventory(item string, count int) {
	i.Lock()
	defer i.Unlock()
	i.items[item] += count
}

func (i *inventory) getInventory(item string) int {
	i.Lock()
	defer i.Unlock()
	return i.items[item]
}

func (i *inventory) printInventory() {
	i.Lock()
	defer i.Unlock()
	fmt.Println("Current Inventory:")
	for item, count := range i.items {
		fmt.Printf("%s: %d\n", item, count)
	}
}

func simulateTransactions(inv *inventory) {
	items := []string{"Apples", "Bananas", "Oranges", "Pineapples", "Mangoes"}
	for {
		item := items[rand.Intn(len(items))]
		count := rand.Intn(10) - 5
		inv.updateInventory(item, count)
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	}
}

func main() {
	inv := &inventory{items: make(map[string]int)}
	inv.updateInventory("Apples", 100)
	inv.updateInventory("Bananas", 80)
	inv.updateInventory("Oranges", 60)
	inv.updateInventory("Pineapples", 40)
	inv.updateInventory("Mangoes", 20)
	inv.printInventory()

	go simulateTransactions(inv)
	go simulateTransactions(inv)
	go simulateTransactions(inv)

	time.Sleep(30 * time.Second)
	inv.printInventory()
}
