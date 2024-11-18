package main

import (
	"context"
	"fmt"
	"log"
	"mongo-go-example/config"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Product struct {
	ID       string  `bson:"_id,omitempty"`
	Name     string  `bson:"name"`
	Quantity int     `bson:"quantity"`
	Price    float32 `bson:"price"`
}

func main() {
	// Connect to the MongoDB database
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(config.MongoDBURI))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(context.TODO())

	// Error handling and graceful termination.
	defer func() {
		if err := recover(); err != nil {
			log.Printf("Panic: %v\n", err)
		}
		client.Disconnect(context.TODO())
	}()

	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()

	database := client.Database(config.DatabaseName)
	collection := database.Collection(config.ProductsCollection)

	// Retrieve data using a Find operation
	findOptions := options.Find().SetLimit(10).SetProjection(&bson.D{{"name", 1}, {"quantity", 1}})
	cursor, err := collection.Find(ctx, bson.D{}, findOptions)
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(ctx)

	var products []Product
	for cursor.Next(ctx) {
		var product Product
		if err := cursor.Decode(&product); err != nil {
			log.Fatal(err)
		}
		products = append(products, product)
	}

	if err := cursor.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Retrieved Products:")
	for _, product := range products {
		fmt.Printf("%+v\n", product)
	}
}
