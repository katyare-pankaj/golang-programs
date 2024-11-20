package main

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// RateLimiter implements a simple Token Bucket rate limiter.
type RateLimiter struct {
	bucket chan struct{}
	rate   int // requests per second
}

func NewRateLimiter(rate int) *RateLimiter {
	return &RateLimiter{
		bucket: make(chan struct{}, rate),
		rate:   rate,
	}
}

func (l *RateLimiter) Wait() {
	t := time.Now()
	d := time.Duration(l.rate) * time.Second
	for {
		select {
		case l.bucket <- struct{}{}:
			return
		case <-time.After(time.Until(t.Add(d))):
			l.bucket <- struct{}{}
			t = time.Now()
		}
	}
}

func streamData(client *mongo.Client, limit int, wg *sync.WaitGroup) {
	defer wg.Done()

	db := client.Database("your_database")
	collection := db.Collection("your_collection")

	pipeline := []bson.M{
		bson.M{"$limit": limit}, // Example: Limit to 1000 documents
	}

	cursor, err := collection.Aggregate(context.TODO(), pipeline, options.Aggregate().SetCursorType(mongo.CursorTypeStream))
	if err != nil {
		log.Fatalf("Error streaming data: %v", err)
	}
	defer cursor.Close(context.TODO())

	rateLimiter := NewRateLimiter(5) // Example: 5 requests per second

	var i int
	for cursor.Next(context.TODO()) {
		i++
		rateLimiter.Wait()

		var doc bson.M
		err := cursor.Decode(&doc)
		if err != nil {
			log.Fatalf("Error decoding document: %v", err)
		}

		// Process document here
		fmt.Printf("Document %d: %v\n", i, doc)
	}

	if err := cursor.Err(); err != nil {
		log.Fatalf("Error iterating cursor: %v", err)
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017/")
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatalf("Error connecting to MongoDB: %v", err)
	}
	defer client.Disconnect(ctx)

	var wg sync.WaitGroup
	wg.Add(1)
	streamData(client, 1000, &wg) // Example: Stream 1000 documents
	wg.Wait()
}
