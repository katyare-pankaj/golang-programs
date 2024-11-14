package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

// Simulate a slow and expensive database query
func getProductFromDB(ctx context.Context, productID string) (*Product, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	case <-time.After(500 * time.Millisecond): // Simulate 500ms query time
	}

	return &Product{
		ID:    productID,
		Name:  fmt.Sprintf("Product %s", productID),
		Price: 19.99,
	}, nil
}

type Product struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

// In-memory cache (can be replaced with a more robust solution like Redis)
var cache = sync.Map{}

// Cache the product data for faster retrieval
func getProductFromCache(productID string) *Product {
	if v, ok := cache.Load(productID); ok {
		return v.(*Product)
	}
	return nil
}

func cacheProduct(product *Product) {
	cache.Store(product.ID, product)
}

func productHandler(w http.ResponseWriter, r *http.Request) {
	productID := r.URL.Query().Get(":product_id")

	ctx, cancel := context.WithTimeout(r.Context(), 1*time.Second)
	defer cancel()

	product := getProductFromCache(productID)
	if product == nil {
		log.Printf("Product %s not in cache, fetching from DB...", productID)
		product, err := getProductFromDB(ctx, productID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		cacheProduct(product)
	}

	fmt.Fprintf(w, "{\"product\":%s}", product)
}

func main() {
	http.HandleFunc("/api/products/{product_id}", productHandler)

	fmt.Println("Server running on http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Error starting server:", err)
	}
}
