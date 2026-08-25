package repository

import (
	"mvc-booklibrary/model"
)

// InMemoryRepository stores books in a slice in memory.
type InMemoryRepository struct {
	books []model.Book
}

// NewInMemoryRepository creates a new in-memory repository.
func NewInMemoryRepository() model.Repository {
	return &InMemoryRepository{
		books: make([]model.Book, 0),
	}
}

func (r *InMemoryRepository) Add(book model.Book) {
	r.books = append(r.books, book)
}

func (r *InMemoryRepository) FindAll() []model.Book {
	return r.books
}

func (r *InMemoryRepository) FindByISBN(isbn string) *model.Book {
	for i, book := range r.books {
		if book.ISBN == isbn {
			return &r.books[i]
		}
	}
	return nil
}

func (r *InMemoryRepository) Update(book model.Book) bool {
	for i, b := range r.books {
		if b.ISBN == book.ISBN {
			r.books[i] = book
			return true
		}
	}
	return false
}

func (r *InMemoryRepository) Remove(isbn string) bool {
	for i, book := range r.books {
		if book.ISBN == isbn {
			r.books = append(r.books[:i], r.books[i+1:]...)
			return true
		}
	}
	return false
}
