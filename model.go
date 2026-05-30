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
