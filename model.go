package main

var library []Book

type Book struct {
	ISBN          string
	Title         string
	Author        string
	PublishedYear int
	Available     bool
}

func AddBook(book Book) {
	library = append(library, book)
}

func FindAllBooks() []Book {
	return library
}
