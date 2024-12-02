package main

import (
	"fmt"
	"sync"
	"time"
)

// Book represents a library book.
type Book struct {
	title  string
	author string
	copies int
	mu     sync.Mutex
}

// Library holds a collection of books.
type Library struct {
	books map[string]*Book
	mu    sync.Mutex
}

// NewLibrary initializes a new library.
func NewLibrary() *Library {
	return &Library{books: make(map[string]*Book)}
}

// AddBook adds a new book to the library.
func (lib *Library) AddBook(title, author string, copies int) {
	lib.mu.Lock()
	defer lib.mu.Unlock()

	if _, exists := lib.books[title]; !exists {
		lib.books[title] = &Book{title: title, author: author, copies: copies}
	}
}

// BorrowBook allows a user to borrow a book if copies are available.
func (lib *Library) BorrowBook(title string) bool {
	lib.mu.Lock()
	book, exists := lib.books[title]
	lib.mu.Unlock()
	if !exists {
		return false
	}

	book.mu.Lock()
	defer book.mu.Unlock()
	if book.copies > 0 {
		book.copies--
		return true
	}
	return false
}

// ReturnBook allows a user to return a borrowed book.
func (lib *Library) ReturnBook(title string) bool {
	lib.mu.Lock()
	book, exists := lib.books[title]
	lib.mu.Unlock()
	if !exists {
		return false
	}

	book.mu.Lock()
	defer book.mu.Unlock()
	book.copies++
	return true
}

// PrintLibraryStatus prints the status of all books in the library.
func (lib *Library) PrintLibraryStatus() {
	lib.mu.Lock()
	defer lib.mu.Unlock()

	fmt.Println("Library Status:")
	for _, book := range lib.books {
		fmt.Printf("Title: %s, Author: %s, Available Copies: %d\n", book.title, book.author, book.copies)
	}
}

func main() {
	library := NewLibrary()

	// Add books to the library
	library.AddBook("The Go Programming Language", "Alan A. A. Donovan", 3)
	library.AddBook("Concurrent Programming in Go", "Katherine Cox-Buday", 2)

	// Simulate multiple users borrowing and returning books
	var wg sync.WaitGroup

	borrowAndReturn := func(userID, title string) {
		defer wg.Done()
		if library.BorrowBook(title) {
			fmt.Printf("User %s: Successfully borrowed \"%s\"\n", userID, title)
			time.Sleep(100 * time.Millisecond) // Simulate reading time
			if library.ReturnBook(title) {
				fmt.Printf("User %s: Successfully returned \"%s\"\n", userID, title)
			}
		} else {
			fmt.Printf("User %s: Failed to borrow \"%s\", no copies available\n", userID, title)
		}
	}

	// Launch multiple goroutines to simulate concurrent book borrowing/returning
	userIDs := []string{"User1", "User2", "User3", "User4", "User5"}
	for i, userID := range userIDs {
		wg.Add(1)
		go borrowAndReturn(userID, "The Go Programming Language")
		if i%2 == 0 { // Every other user tries to borrow a different book
			wg.Add(1)
			go borrowAndReturn(userID, "Concurrent Programming in Go")
		}
	}

	wg.Wait()

	// Print final status of the library
	library.PrintLibraryStatus()
}
