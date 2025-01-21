package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Book struct {
	ID     int
	Title  string
	Author string
}

func main() {
	books := generateBooks(1000)
	start := time.Now()
	MigrateBooksParallel(books, 10) //migrate 1000 books using 10 goroutines
	elapsed := time.Since(start)
	fmt.Println("Migration completed in", elapsed)
}

func generateBooks(n int) []*Book {
	books := make([]*Book, n)
	for i := 0; i < n; i++ {
		books[i] = &Book{ID: i + 1, Title: fmt.Sprintf("Book %d", i+1), Author: fmt.Sprintf("Author %d", i+1)}
	}
	return books
}

func MigrateBooksParallel(books []*Book, concurrency int) {
	var wg sync.WaitGroup
	wg.Add(len(books))

	migrate := func(start, end int, wg *sync.WaitGroup) {
		defer wg.Done()
		for i := start; i < end; i++ {
			err := migrateBook(books[i])
			if err != nil {
				fmt.Printf("Error migrating book %d: %v\n", books[i].ID, err)
			}
		}
	}

	// Partition the slice into chunks and start goroutines to migrate each chunk
	chunkSize := len(books)/concurrency + 1
	for i := 0; i < len(books); i += chunkSize {
		end := i + chunkSize
		if end > len(books) {
			end = len(books)
		}
		go migrate(i, end, &wg)
	}

	wg.Wait()
}

func migrateBook(book *Book) error {
	// Simulate database migration delay
	time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)

	// Uncomment the following line to add an artificial error
	// if book.ID % 13 == 0 {
	//  return fmt.Errorf("Artificial error migrating book %d", book.ID)
	// }

	fmt.Printf("Migrated book %d: %s by %s\n", book.ID, book.Title, book.Author)
	return nil
}
