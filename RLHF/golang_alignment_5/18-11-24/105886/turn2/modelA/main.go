package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var (
	db   *sql.DB
	pool *sync.Pool
)

func init() {
	dbConfig := "user:password@tcp(localhost:3306)/your_database?parseTime=True&loc=Local"
	db, err := sql.Open("mysql", dbConfig)
	if err != nil {
		log.Fatal("Error connecting to database:", err)
	}

	pool = &sync.Pool{
		New: func() interface{} {
			conn, err := db.Conn(context.Background())
			if err != nil {
				log.Fatal("Error creating new connection:", err)
			}
			return conn
		},
	}
}

func queryExample() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	conn := pool.Get().(sql.Conn)
	defer pool.Put(conn)

	rows, err := conn.QueryContext(ctx, "SELECT * FROM your_table")
	if err != nil {
		log.Error("Error executing query:", err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var name string
		err := rows.Scan(&id, &name)
		if err != nil {
			log.Error("Error scanning row:", err)
			continue
		}
		fmt.Printf("ID: %d, Name: %s\n", id, name)
	}

	if err := rows.Err(); err != nil {
		log.Error("Error after rows.Next():", err)
	}
}

func main() {
	for {
		queryExample()
		time.Sleep(1 * time.Second)
	}
}
