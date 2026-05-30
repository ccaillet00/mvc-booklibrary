package main

import (
	"encoding/json"
	"os"
)

type JSONRepository struct {
	filepath string
}

func NewJSONRepository(filepath string) *JSONRepository {
	return &JSONRepository{filepath: filepath}
}

func (r *JSONRepository) Add(book Book) {
	books := r.load()
	books = append(books, book)
	r.save(books)
}

func (r *JSONRepository) FindAll() []Book {
	return r.load()
}

func (r *JSONRepository) FindByISBN(isbn string) *Book {
	books := r.load()
	for i, book := range books {
		if book.ISBN == isbn {
			return &books[i]
		}
	}
	return nil
}

func (r *JSONRepository) Update(book Book) bool {
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

func (r *JSONRepository) load() []Book {
	data, err := os.ReadFile(r.filepath)
	if err != nil {
		return []Book{} // Datei existiert nicht? Leeres Slice!
	}
	var books []Book
	json.Unmarshal(data, &books) // JSON → Go
	return books
}

func (r *JSONRepository) save(books []Book) {
	data, _ := json.MarshalIndent(books, "", "  ") // Go → JSON (schön formatiert)
	os.WriteFile(r.filepath, data, 0644)           // In Datei schreiben
}
