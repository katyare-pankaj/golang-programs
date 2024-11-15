package main

import (
	"fmt"

	"github.com/boltdb/bolt"
)

var db *bolt.DB // Declare db globally

func initWriteDB() {
	var err error
	// Open the database (or create it if it doesn't exist)
	db, err = bolt.Open("write.db", 0600, nil) // Use = to assign to the global db
	if err != nil {
		panic(err)
	}

	// Initialize the database by creating the "users" bucket if it doesn't exist
	err = db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("users"))
		return err
	})
	if err != nil {
		panic(err)
	}
}

func writeUser(user string) {
	// Insert a user into the "users" bucket
	err := db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("users"))
		if bucket == nil {
			return fmt.Errorf("Bucket not found")
		}
		err := bucket.Put([]byte(user), []byte{})
		return err
	})
	if err != nil {
		fmt.Println("Error writing user:", err)
	}
}

func readUsers() {
	// Read the users from the "users" bucket
	err := db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("users"))
		if bucket == nil {
			return fmt.Errorf("Bucket not found")
		}

		// Iterate over all the keys in the "users" bucket
		return bucket.ForEach(func(k, v []byte) error {
			fmt.Printf("User: %s\n", k) // Print the user key
			return nil
		})
	})
	if err != nil {
		fmt.Println("Error reading users:", err)
	}
}

func main() {
	// Initialize and connect to the database
	initWriteDB()

	// Write users to the database
	writeUser("alice")
	writeUser("bob")
	writeUser("chris")

	// Read and print users
	readUsers()
}
