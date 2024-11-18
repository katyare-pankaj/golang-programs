package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Product struct {
	ID    bson.ObjectId `bson:"_id,omitempty"`
	Name  string        `bson:"name"`
	Price float64       `bson:"price"`
}

func fetchProducts(db *mgo.Database) ([]Product, error) {
	c := db.C("products")

	// Set a query timeout to prevent infinite waits
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var products []Product
	err := c.Find(nil).Limit(100).All(&products) // Limit results to 100 for efficient memory usage
	if err != nil {
		return nil, fmt.Errorf("error fetching products: %w", err)
	}

	return products, nil
}

func main() {
	session, err := mgo.Dial("mongodb://localhost:27017/retail_analytics")
	if err != nil {
		log.Fatalf("error connecting to MongoDB: %w", err)
	}
	defer session.Close()

	db := session.DB("")
	db.SetReadPreference(&mgo.ReadPreference{Primary: true}) // Opt for primary read preference

	products, err := fetchProducts(db)
	if err != nil {
		log.Fatalf("error fetching products: %w", err)
	}

	for _, product := range products {
		fmt.Println(product)
	}
}
