package model

// Book represents a book in the library.
type Book struct {
	ISBN          string
	Title         string
	Author        string
	PublishedYear int
	Borrowed      bool
}

// Repository defines the interface for book persistence.
type Repository interface {
	Add(book Book)
	FindAll() []Book
	FindByISBN(isbn string) *Book
	Update(book Book) bool
	Remove(isbn string) bool
}

// Model contains the business logic for the book library.
type Model struct {
	repo Repository
}

// NewModel creates a new Model with the given repository.
func NewModel(repo Repository) *Model {
	return &Model{repo: repo}
}

// AddBook adds a book to the library.
func (m *Model) AddBook(b Book) {
	m.repo.Add(b)
}

// FindAllBooks returns all books in the library.
func (m *Model) FindAllBooks() []Book {
	return m.repo.FindAll()
}

func (m *Model) FindBook(isbn string) *Book {
	return m.repo.FindByISBN(isbn)
}

// CheckBookAvailability checks if a book exists and is not borrowed.
func (m *Model) CheckBookAvailability(isbn string) bool {
	book := m.repo.FindByISBN(isbn)
	return book != nil && !book.Borrowed
}

// LendBook marks a book as borrowed.
func (m *Model) LendBook(isbn string) *Book {
	book := m.repo.FindByISBN(isbn)
	if book != nil {
		book.Borrowed = true
		m.repo.Update(*book)
		return book
	}
	return nil
}

// ReturnBook marks a book as returned.
func (m *Model) ReturnBook(isbn string) *Book {
	book := m.repo.FindByISBN(isbn)
	if book != nil && book.Borrowed {
		book.Borrowed = false
		m.repo.Update(*book)
		return book
	}
	return nil
}

// RemoveBook removes a book from the library.
func (m *Model) RemoveBook(isbn string) bool {
	return m.repo.Remove(isbn)
}

// AvailableBooks returns all books that are not borrowed.
func (m *Model) AvailableBooks() []Book {
	var available []Book
	for _, b := range m.repo.FindAll() {
		if !b.Borrowed {
			available = append(available, b)
		}
	}
	return available
}

// BorrowedBooks returns all books that are currently borrowed.
func (m *Model) BorrowedBooks() []Book {
	var borrowed []Book
	for _, b := range m.repo.FindAll() {
		if b.Borrowed {
			borrowed = append(borrowed, b)
		}
	}
	return borrowed
}

// FindDuplicateISBN checks if a book with the given ISBN appears more than once.
func (m *Model) FindDuplicateISBN(isbn string) bool {
	var count int
	for _, b := range m.repo.FindAll() {
		if b.ISBN == isbn {
			count++
		}
	}
	return count > 1
}
