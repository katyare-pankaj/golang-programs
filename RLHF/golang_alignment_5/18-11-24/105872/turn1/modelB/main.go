package main

import (
	"fmt"
	"log"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Product represents a product document in the MongoDB database
type Product struct {
	ID    bson.ObjectId `bson:"_id,omitempty"`
	Name  string        `bson:"name"`
	Price float64       `bson:"price"`
}

func fetchProducts(db *mgo.Database) ([]Product, error) {
	c := db.C("products")

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

	products, err := fetchProducts(db)
	if err != nil {
		log.Fatalf("error fetching products: %v", err)
	}

	for _, product := range products {
		fmt.Println(product)
	}
}
