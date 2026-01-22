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

func CheckBookAvailability(isbn string) bool {
	for _, Book := range library {
		if Book.ISBN == isbn && Book.Available {
			return true
		}
	}
	return false
}

func LendBook(isbn string) *Book {
	for i, Book := range library {
		if Book.ISBN == isbn && Book.Available {
			library[i].Available = false
			return &library[i]
		}
	}
	return nil
}

func ReturnBook(isbn string) *Book {
	for i, Book := range library {
		if Book.ISBN == isbn && !Book.Available {
			library[i].Available = true
			return &library[i]
		}
	}
	return nil
}

func RemoveBook(isbn string) bool {
	for i, Book := range library {
		if Book.ISBN == isbn {
			library = append(library[:i], library[i+1:]...)
			return true
		}
	}
	return false
}
