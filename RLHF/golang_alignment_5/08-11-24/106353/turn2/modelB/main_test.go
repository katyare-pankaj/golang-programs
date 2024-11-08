package main

import (
	"testing"
)

func TestInventoryAddItem(t *testing.T) {
	inv := &Inventory{Items: make(map[string]int)}

	inv.AddItem("apple", 5)
	inv.AddItem("banana", 3)

	if inv.GetItemCount("apple") != 5 {
		t.Errorf("Item count for apple should be 5, got: %d", inv.GetItemCount("apple"))
	}
	if inv.GetItemCount("banana") != 3 {
		t.Errorf("Item count for banana should be 3, got: %d", inv.GetItemCount("banana"))
	}
}
