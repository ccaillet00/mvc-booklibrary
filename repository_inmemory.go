package main

type InMemoryRepository struct {
	books []Book
}

func (r *InMemoryRepository) Add(book Book) {
	r.books = append(r.books, book)
}

func (r *InMemoryRepository) FindAll() []Book {
	return r.books
}

func (r *InMemoryRepository) FindByISBN(isbn string) *Book {
	for i, book := range r.books {
		if book.ISBN == isbn {
			return &r.books[i]
		}
	}
	return nil
}

func (r *InMemoryRepository) Update(book Book) bool {
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

func (r *InMemoryRepository) RemoveAll() {

}

func NewInMemoryRepository() Repository {
	return &InMemoryRepository{
		books: make([]Book, 0),
	}
}
