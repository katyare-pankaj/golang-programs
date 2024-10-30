package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/patrickmn/go-cache"
)

// Product represents a product in the e-commerce store
type Product struct {
	ID          int
	Name        string
	Description string
	Price       float32
	Category    string
}

// InMemoryCacheStore uses in-memory caching to store and retrieve products
type InMemoryCacheStore struct {
	cache *cache.Cache
	mu    sync.Mutex
}

// NewInMemoryCacheStore creates a new InMemoryCacheStore
func NewInMemoryCacheStore() *InMemoryCacheStore {
	return &InMemoryCacheStore{
		cache: cache.New(5*time.Minute, 10*time.Minute), // Cache items for 5 minutes, clean up every 10 minutes
	}
}

// StoreProduct stores a product in the in-memory cache
func (ims *InMemoryCacheStore) StoreProduct(product *Product) {
	ims.mu.Lock()
	defer ims.mu.Unlock()
	ims.cache.Set(fmt.Sprint(product.ID), product, cache.DefaultExpiration)
}

// GetProduct retrieves a product from the in-memory cache based on its ID
func (ims *InMemoryCacheStore) GetProduct(id int) (*Product, bool) {
	ims.mu.Lock()
	defer ims.mu.Unlock()
	val, found := ims.cache.Get(fmt.Sprint(id))
	if found {
		return val.(*Product), true
	}
	return nil, false
}

// Simulate database query to fetch product by ID
func fetchProductFromDB(id int) *Product {
	// Simulate delay for database query
	time.Sleep(time.Millisecond * 500)

	switch id {
	case 1:
		return &Product{ID: 1, Name: "Android Phone", Description: "Flagship smartphone", Price: 799.99, Category: "Electronics"}
	case 2:
		return &Product{ID: 2, Name: "Women's Jeans", Description: "Stylish blue jeans", Price: 49.99, Category: "Clothing"}
	case 3:
		return &Product{ID: 3, Name: "Book: The Alchemist", Description: "A novel by Paulo Coelho", Price: 12.99, Category: "Books"}
	default:
		return nil
	}
}

// GetProductWithCache retrieves a product using the in-memory cache. If the product is not in the cache, it fetches it from the database and stores it in the cache.
func GetProductWithCache(store *InMemoryCacheStore, id int) *Product {
	product, found := store.GetProduct(id)
	if found {
		fmt.Println("Product found in cache.")
		return product
	}

	fmt.Println("Product not found in cache, fetching from database...")
	product = fetchProductFromDB(id)
	if product != nil {
		store.StoreProduct(product)
	}
	return product
}

func main() {
	store := NewInMemoryCacheStore()

	// Perform multiple searches for the same product IDs to observe caching behavior
	productIDs := []int{1, 2, 3, 1, 2}

	for _, id := range productIDs {
		product := GetProductWithCache(store, id)
		if product != nil {
			fmt.Printf("Product Found: %+v\n", product)
		} else {
			fmt.Printf("Product with ID %d not found.\n", id)
		}
	}
}
