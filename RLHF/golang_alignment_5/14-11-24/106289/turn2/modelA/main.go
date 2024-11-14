package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq" // PostgreSQL driver
)

// Product represents a product in the catalog
type Product struct {
	ID        int     `json:"id"`
	Name      string  `json:"name"`
	Price     float32 `json:"price"`
	Inventory int     `json:"inventory"`
}

var cache = make(map[int]Product)
var cacheMu sync.RWMutex

func getProductFromDB(id int) Product {
	// Simulate database query
	time.Sleep(time.Duration(200) * time.Millisecond)

	// Replace this with actual DB query using `database/sql`
	prod := Product{
		ID:        id,
		Name:      fmt.Sprintf("Product %d", id),
		Price:     19.99,
		Inventory: 100,
	}

	return prod
}

func getProductFromCache(id int) Product {
	cacheMu.RLock()
	defer cacheMu.RUnlock()

	if product, ok := cache[id]; ok {
		return product
	}

	return Product{}
}

func storeProductInCache(id int, product Product) {
	cacheMu.Lock()
	defer cacheMu.Unlock()

	cache[id] = product
}

func getProductHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	prod := getProductFromCache(id)
	if prod.ID == 0 {
		prod = getProductFromDB(id)
		if prod.ID != 0 {
			storeProductInCache(id, prod)
		}
	}

	if prod.ID == 0 {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(prod)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/products/{id}", getProductHandler).Methods("GET")

	fmt.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
