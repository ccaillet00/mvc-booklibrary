package main

type Book struct {
	ISBN          string
	Title         string
	Author        string
	PublishedYear int
	Available     bool
}

type Repository interface {
	Add(book Book)
	FindAll() []Book
	FindByISBN(isbn string) *Book
	Update(book Book) bool
	Remove(isbn string) bool
}

var repo Repository = NewInMemoryRepository()

func AddBook(book Book) {
	repo.Add(book)
}

func FindAllBooks() []Book {
	return repo.FindAll()
}

func CheckBookAvailability(isbn string) bool {
	book := repo.FindByISBN(isbn)
	return book != nil && book.Available
}

func LendBook(isbn string) *Book {
	book := repo.FindByISBN(isbn)
	if book != nil && book.Available {
		book.Available = false
		repo.Update(*book)
		return book
	}
	return nil
}

func ReturnBook(isbn string) *Book {
	book := repo.FindByISBN(isbn)
	if book != nil && !book.Available {
		book.Available = true
		repo.Update(*book)
		return book
	}
	return nil
}

func RemoveBook(isbn string) bool {
	return repo.Remove(isbn)
}
