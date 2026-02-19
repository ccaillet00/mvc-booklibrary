package main

type InMemoryRepository struct {
	library []Book
}

func (r *InMemoryRepository) Add(book Book) {
	r.library = append(r.library, book)
}

func (r *InMemoryRepository) FindAll() []Book {
	return r.library
}

func (r *InMemoryRepository) FindByISBN(isbn string) *Book {
	for i, Book := range r.library {
		if Book.ISBN == isbn {
			return &r.library[i]
		}
	}
	return nil
}

func (r *InMemoryRepository) Update(book Book) bool {
	for i, Book := range r.library {
		if Book.ISBN == book.ISBN {
			r.library[i] = book
			return true
		}
	}
	return false
}

func (r *InMemoryRepository) Remove(isbn string) bool {
	for i, book := range r.library {
		if book.ISBN == isbn {
			r.library = append(r.library[:i], r.library[i+1:]...)
			return true
		}
	}
	return false
}

func NewInMemoryRepository() *InMemoryRepository {
	return &InMemoryRepository{
		library: make([]Book, 0),
	}
}
