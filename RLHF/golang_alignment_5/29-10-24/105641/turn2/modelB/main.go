package main  
import (  
	"fmt"
	"math"
	"math/rand"
	"sort"
	"time"
)

// Order represents a supply chain order
type Order struct {
	ID        int
	ItemID    int
	Quantity  int
	ProcessingTime int // Time taken to process the order (in days)
}

// Inventory represents the inventory status of an item
type Inventory struct {
	ItemID    int
	CurrentStock int
	MinimumStock int
	MaximumStock int
	DemandForecast int // Forecasted demand for the item (in units)
}

// DemandHistory stores demand data for an item over time
type DemandHistory struct{
	ItemID int
	Demand []int
}

// Supplier represents a supplier for an item
type Supplier struct{
	ItemID int
	LeadTime int // Time taken to receive an order from the supplier (in days)
}

//CalculateReorderLevel calculates the reorder level for an item
func CalculateReorderLevel(inventory *Inventory) int {
	return int(math.Ceil(float64(inventory.DemandForecast) * (float64(inventory.MaximumStock - inventory.MinimumStock) / 100))) + inventory.MinimumStock
}

//UpdateInventory updates the inventory after processing an order
func UpdateInventory(inventory *Inventory, order *Order) {
	inventory.CurrentStock -= order.Quantity
}

//ProcessOrder processes an order and updates inventory
func ProcessOrder(inventoryMap map[int]*Inventory, order *Order) {
	// Find the inventory for the item being ordered
	inventory, ok := inventoryMap[order.ItemID]
	if !ok {
		fmt.Printf("Error: Item ID %d not found in inventory.\n", order.ItemID)
		return
	}

	// Update inventory after processing the order
	UpdateInventory(inventory, order)

	// Check if reorder level needs to be updated
	if inventory.CurrentStock <= inventory.ReorderLevel {
		fmt.Printf("Reorder level reached for Item ID %d. Placing a reorder.\n", inventory.ItemID)
		// Place a reorder here (not implemented in this example)
	}
}

// PredictDemand uses simple statistical methods to predict demand for an item
func PredictDemand(demandHistory []DemandHistory, itemID int, period int) int {
	// We'll use a simple moving average for prediction in this example
	sum := 0
	count := 0

	// Find the demand history for the specified item
	for _, history := range demandHistory {
		if history.ItemID == itemID {
			for _, demand := range history.Demand {
				if count >= period {
					break
				}
				sum += demand
				count++
			}
			break
		}
	}

	if count == 0 {
		return 0 // Return 0 if no demand history is found
	}

	return int(math.Round(float64(sum) / float64(count)))
}

const (
	numItems = 5
	maxDemand = 20
	numOrders = 20
	reorderPercentage = 20 // Percentage of maximum stock as reorder level
)

func main() {
	rand.Seed(time.Now().UnixNano())

	// Initialize inventory
	inventoryMap := make(map[int]*Inventory)
	for i := 0; i < numItems; i++ {
		minStock := rand.Intn(10) + 1
		maxStock := minStock + rand.Intn(20) + 1
		inventory