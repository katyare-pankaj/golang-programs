package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Product represents a product document in the MongoDB database
type Product struct {
	ID       bson.ObjectId `bson:"_id,omitempty"`
	Name     string        `bson:"name"`
	Price    float64       `bson:"price"`
	Quantity int           `bson:"quantity"`
}

func fetchProductsWithEfficientHandling(db *mgo.Database) ([]Product, error) {
	c := db.C("products")
	_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var products []Product
	err := c.Find(nil).All(&products)
	if err != nil {
		return nil, fmt.Errorf("error fetching products: %w", err)
	}

	return products, nil
}

func main() {
	session, err := mgo.Dial("mongodb://localhost:27017/retail_analytics")
	if err != nil {
		log.Fatalf("error connecting to MongoDB: %v", err)
	}
	defer session.Close()
	db := session.DB("")
	defer db.Session.Close() // Ensure session is closed properly

	products, err := fetchProductsWithEfficientHandling(db)
	if err != nil {
		log.Printf("Error fetching products: %v", err)
		return
	}

	for _, product := range products {
		fmt.Println(product)
	}
}
