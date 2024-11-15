package main

import (
	"fmt"

	"github.com/boltdb/bolt"
)

var db *bolt.DB

func initWriteDB() {
	db, err := bolt.Open("write.db", 0600, nil)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists("users")
		return err
	})
	if err != nil {
		panic(err)
	}
}

func writeUser(user string) {
	err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket("users")
		err := b.Put([]byte(user), []byte{})
		return err
	})
	if err != nil {
		fmt.Println("Error writing user:", err)
	}
}

func main() {
	initWriteDB()
	writeUser("alice")
	writeUser("bob")
	writeUser("chris")
}
