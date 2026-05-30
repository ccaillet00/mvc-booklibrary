package main

type Repository interface {
	Add(book Book)
	FindAll() []Book
	FindByISBN(isbn string) *Book
	Update(book Book) bool
	Remove(isbn string) bool
}
