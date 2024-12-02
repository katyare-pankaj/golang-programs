package main

import (
	"fmt"
	"sync"
)

type Library struct {
	books     []string
	lock      sync.Mutex
	condition sync.Cond
}

func NewLibrary(books []string) *Library {
	return &Library{books: books, condition: sync.NewCond(&sync.Mutex{})}
}

// Borrow borrows a book from the library
func (lib *Library) Borrow(user string) (string, error) {
	lib.condition.L.Lock()
	defer lib.condition.L.Unlock()

	for i, book := range lib.books {
		if lib.books[i] != "" {
			lib.books[i] = "" // Mark book as borrowed
			lib.condition.Signal()
			return book, nil
		}
	}

	return "", fmt.Errorf("%s: No books available for borrowing", user)
}

// Return returns a book to the library
func (lib *Library) Return(user string, book string) error {
	lib.condition.L.Lock()
	defer lib.condition.L.Unlock()

	for i, libBook := range lib.books {
		if libBook == "" {
			lib.books[i] = book
			lib.condition.Signal()
			return nil
		}
	}

	return fmt.Errorf("%s: Cannot return book, library is full", user)
}

func main() {
	library := NewLibrary([]string{"Book A", "Book B", "Book C", "Book D"})
	numUsers := 5
	numActions := 5

	wg := sync.WaitGroup{}

	for user := 1; user <= numUsers; user++ {
		wg.Add(1)
		go func(user int) {
			defer wg.Done()
			for action := 1; action <= numActions; action++ {
				doAction(library, fmt.Sprintf("User %d", user))
			}
		}(user)
	}

	wg.Wait()
}

func doAction(lib *Library, user string) {
	fmt.Printf("%s: Waiting for action...\n", user)
	lib.condition.L.Lock()
	defer lib.condition.L.Unlock()

	for {
		if len(lib.books) > 0 {
			break
		}
		lib.condition.Wait()
	}

	action := "Borrow"
	if rand()%2 == 1 {
		action = "Return"
	}

	fmt.Printf("%s: %s a book...\n", user, action)
	if action == "Borrow" {
		book, err := lib.Borrow(user)
		if err != nil {
			fmt.Printf("%s: Failed to borrow book: %v\n", user, err)
		} else {
			fmt.Printf("%s: Borrowed book: %s\n", user, book)
		}
	} else {
		book := "Book A" // Choose a book to return
		err := lib.Return(user, book)
		if err != nil {
			fmt.Printf("%s: Failed to return book: %v\n", user, err)
		} else {
			fmt.Printf("%s: Returned book: %s\n", user, book)
		}
	}
}
