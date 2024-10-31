// catalog/module.go

package catalog

import (
	"fmt"
	"go-programs/RLHF/golang_alignment_5/30-10-24/105724/turn2/modelA/library"
)

// DisplayBooks displays a list of books in the catalog.
func DisplayBooks(books []library.Book) {
	fmt.Println("Books in the Catalog:")
	for _, book := range books {
		fmt.Printf("%s by %s (%s)\n", book.Title, book.Author, book.ISBN)
	}
}

// SearchBookByISBN searches for a book by its ISBN in the catalog.
func SearchBookByISBN(books []library.Book, isbn string) *library.Book {
	for _, book := range books {
		if book.ISBN == isbn {
			return &book
		}
	}
	return nil
}
