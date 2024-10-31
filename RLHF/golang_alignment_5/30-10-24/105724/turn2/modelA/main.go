// main.go

package main

import (
	"fmt"
	"go-programs/RLHF/golang_alignment_5/30-10-24/105724/turn2/modelA/catalog"
	"go-programs/RLHF/golang_alignment_5/30-10-24/105724/turn2/modelA/library"
)

func main() {
	// Create a new library
	lib := library.Library{}

	// Add some books to the library
	lib.AddBook(library.Book{Title: "The Alchemist", Author: "Paulo Coelho", ISBN: "9780062564811"})
	lib.AddBook(library.Book{Title: "To Kill a Mockingbird", Author: "Harper Lee", ISBN: "9780060935467"})

	// Display the books in the catalog
	catalog.DisplayBooks(lib.GetBooks())

	// Search for a book by ISBN
	isbnToSearch := "9780060935467"
	foundBook := catalog.SearchBookByISBN(lib.GetBooks(), isbnToSearch)
	if foundBook != nil {
		fmt.Printf("Found book: %s by %s\n", foundBook.Title, foundBook.Author)
	} else {
		fmt.Printf("Book with ISBN '%s' not found.\n", isbnToSearch)
	}
}
