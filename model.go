package main

type Model struct {
	repo Repository
}

func NewModel(repo Repository) Model {
	return Model{repo}
}

type Book struct {
	ISBN          string
	Title         string
	Author        string
	PublishedYear int
	Borrowed      bool
}

func (m *Model) AddBook(b Book) {
	m.repo.Add(b)
}

func (m *Model) FindAllBooks() []Book {
	return m.repo.FindAll()
}

func (m *Model) CheckBookAvailability(isbn string) bool {
	book := m.repo.FindByISBN(isbn)
	return book != nil && !book.Borrowed
}

func (m *Model) LendBook(isbn string) *Book {
	book := m.repo.FindByISBN(isbn)
	if book != nil {
		book.Borrowed = true
		m.repo.Update(*book)
		return book
	}
	return nil
}

func (m *Model) ReturnBook(isbn string) *Book {
	book := m.repo.FindByISBN(isbn)
	if book != nil && book.Borrowed {
		book.Borrowed = false
		m.repo.Update(*book)
		return book
	}
	return nil
}

func (m *Model) RemoveBook(isbn string) bool {
	return m.repo.Remove(isbn)
}

func (m *Model) AvailableBooks() []Book {
	var available []Book
	for _, b := range m.repo.FindAll() {
		if !b.Borrowed {
			available = append(available, b)
		}
	}
	return available
}

func (m *Model) BorrowedBooks() []Book {
	var borrowed []Book
	for _, b := range m.repo.FindAll() {
		if b.Borrowed {
			borrowed = append(borrowed, b)
		}
	}
	return borrowed
}

func (m *Model) FindDuplicateISBN(isbn string) bool {
	var count int
	for _, b := range m.repo.FindAll() {
		if b.ISBN == isbn {
			count++
		}
	}
	return count > 1
}
