// library/module.go

package library

// Book represents a book in the library.
type Book struct {
	Title  string
	Author string
	ISBN   string
}

// Library represents a collection of books.
type Library struct {
	Books []Book
}

// AddBook adds a new book to the library.
func (l *Library) AddBook(book Book) {
	l.Books = append(l.Books, book)
}

// GetBooks returns all the books in the library.
func (l *Library) GetBooks() []Book {
	return l.Books
}
