package model

import (
	"testing"
)

// testRepo is a minimal in-memory repository for testing,
// defined here to avoid an import cycle with the repository package.
type testRepo struct {
	books []Book
}

func (r *testRepo) Add(book Book) {
	r.books = append(r.books, book)
}

func (r *testRepo) FindAll() []Book {
	return r.books
}

func (r *testRepo) FindByISBN(isbn string) *Book {
	for i, b := range r.books {
		if b.ISBN == isbn {
			return &r.books[i]
		}
	}
	return nil
}

func (r *testRepo) Update(book Book) bool {
	for i, b := range r.books {
		if b.ISBN == book.ISBN {
			r.books[i] = book
			return true
		}
	}
	return false
}

func (r *testRepo) Remove(isbn string) bool {
	for i, b := range r.books {
		if b.ISBN == isbn {
			r.books = append(r.books[:i], r.books[i+1:]...)
			return true
		}
	}
	return false
}

func newTestRepo() *testRepo {
	return &testRepo{books: make([]Book, 0)}
}

type spyRepository struct {
	Repository
	updateCalls int
}

func (s *spyRepository) Update(book Book) bool {
	s.updateCalls++
	return s.Repository.Update(book)
}

func TestModel_AddBook(t *testing.T) {
	repo := newTestRepo()
	m := NewModel(repo)

	m.AddBook(Book{ISBN: "111", Title: "Clean Code"})

	books := m.FindAllBooks()
	if len(books) != 1 {
		t.Fatalf("erwartet 1 Buch, bekommen %d", len(books))
	}
}

func TestModel_LendBook(t *testing.T) {
	spy := &spyRepository{Repository: newTestRepo()}
	m := NewModel(spy)
	m.AddBook(Book{ISBN: "111", Title: "Clean Code"})

	m.LendBook("111")

	if spy.updateCalls != 1 {
		t.Errorf("LendBook sollte Update 1x aufrufen, war %d", spy.updateCalls)
	}
}

func TestModel_ReturnBook(t *testing.T) {
	m := NewModel(newTestRepo())
	m.AddBook(Book{ISBN: "111", Title: "Clean Code", Borrowed: true})

	book := m.ReturnBook("111")

	if book == nil {
		t.Fatal("ReturnBook gab nil zurück - erwartet das zurückgegebene Buch")
	}
	if book.Borrowed {
		t.Error("Buch sollte nach Rückgabe nicht mehr ausgeliehen sein")
	}
}

func TestModel_CheckBookAvailability(t *testing.T) {
	m := NewModel(newTestRepo())
	m.AddBook(Book{ISBN: "111", Title: "Clean Code"})

	available := m.CheckBookAvailability("111")

	if !available {
		t.Error("Buch sollte verfügbar sein")
	}
}

func TestModel_AvailableBooks(t *testing.T) {
	m := NewModel(newTestRepo())
	m.AddBook(Book{ISBN: "1", Borrowed: true})
	m.AddBook(Book{ISBN: "2", Borrowed: false})
	m.AddBook(Book{ISBN: "3", Borrowed: true})

	available := m.AvailableBooks()

	if len(available) != 1 {
		t.Fatalf("erwartet 1 verfügbares Buch, bekommen %d", len(available))
	}
	if available[0].ISBN != "2" {
		t.Errorf("falsches Buch: erwartete ISBN '2', bekommen '%s'", available[0].ISBN)
	}
}

func TestModel_BorrowedBooks(t *testing.T) {
	m := NewModel(newTestRepo())
	m.AddBook(Book{ISBN: "1", Borrowed: true})
	m.AddBook(Book{ISBN: "2", Borrowed: false})
	m.AddBook(Book{ISBN: "3", Borrowed: true})

	borrowed := m.BorrowedBooks()

	if len(borrowed) != 2 {
		t.Fatalf("erwartet 2 ausgeliehene Bücher, bekommen %d", len(borrowed))
	}
	if borrowed[0].ISBN != "1" || borrowed[1].ISBN != "3" {
		t.Errorf("falsche Bücher: erwartete '1' und '3', bekommen '%s' und '%s'", borrowed[0].ISBN, borrowed[1].ISBN)
	}
}

func TestModel_FindDuplicateISBN(t *testing.T) {
	m := NewModel(newTestRepo())
	m.AddBook(Book{ISBN: "111", Title: "Clean Code"})
	m.AddBook(Book{ISBN: "111", Title: "Clean Code 2"})
	m.AddBook(Book{ISBN: "222", Title: "The Pragmatic Programmer"})

	duplicate := m.FindDuplicateISBN("111")

	if !duplicate {
		t.Error("erwartet, dass eine doppelte ISBN gefunden wird")
	}
}

func TestDuplicateISBN(t *testing.T) {
	m := NewModel(newTestRepo())
	m.AddBook(Book{ISBN: "111", Title: "Clean Code"})
	m.AddBook(Book{ISBN: "111", Title: "Clean Code 2"})
	m.AddBook(Book{ISBN: "222", Title: "The Pragmatic Programmer"})

	duplicate := m.FindDuplicateISBN("111")

	if !duplicate {
		t.Errorf("erwartet, dass ein doppeltes ISBN gefunden wird")
	}
}
