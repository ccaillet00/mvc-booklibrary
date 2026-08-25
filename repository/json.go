package repository

import (
	"encoding/json"
	"os"

	"mvc-booklibrary/model"
)

// JSONRepository persists books to a JSON file.
type JSONRepository struct {
	filepath string
}

// NewJSONRepository creates a new JSON file repository.
func NewJSONRepository(filepath string) model.Repository {
	return &JSONRepository{filepath: filepath}
}

func (r *JSONRepository) Add(book model.Book) {
	books := r.load()
	books = append(books, book)
	r.save(books)
}

func (r *JSONRepository) FindAll() []model.Book {
	return r.load()
}

func (r *JSONRepository) FindByISBN(isbn string) *model.Book {
	books := r.load()
	for i, book := range books {
		if book.ISBN == isbn {
			return &books[i]
		}
	}
	return nil
}

func (r *JSONRepository) Update(book model.Book) bool {
	books := r.load()
	for i, b := range books {
		if b.ISBN == book.ISBN {
			books[i] = book
			r.save(books)
			return true
		}
	}
	return false
}

func (r *JSONRepository) Remove(isbn string) bool {
	books := r.load()
	for i, book := range books {
		if book.ISBN == isbn {
			books = append(books[:i], books[i+1:]...)
			r.save(books)
			return true
		}
	}
	return false
}

func (r *JSONRepository) load() []model.Book {
	data, err := os.ReadFile(r.filepath)
	if err != nil {
		return []model.Book{}
	}
	var books []model.Book
	json.Unmarshal(data, &books)
	return books
}

func (r *JSONRepository) save(books []model.Book) {
	data, _ := json.MarshalIndent(books, "", "  ")
	os.WriteFile(r.filepath, data, 0644)
}
