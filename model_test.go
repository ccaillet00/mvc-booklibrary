package main

import (
	"testing"
)

type spyRepository struct {
	Repository
	updateCalls int
}

func (s *spyRepository) Update(book Book) bool {
	s.updateCalls++
	return s.Repository.Update(book)
}

func TestModel_AddBook(t *testing.T) {
	repo := NewInMemoryRepository()
	model := NewModel(repo)

	model.AddBook(Book{ISBN: "111", Title: "Clean Code"})

	books := model.FindAllBooks()
	if len(books) != 1 {
		t.Fatalf("erwartet 1 Buch, bekommen %d", len(books))
	}
}

func TestModel_LendBook(t *testing.T) {
	spy := &spyRepository{Repository: NewInMemoryRepository()}
	model := NewModel(spy)
	model.AddBook(Book{ISBN: "111", Title: "Clean Code"})

	model.LendBook("111")

	if spy.updateCalls != 1 {
		t.Errorf("LendBook sollte Update 1x aufrufen, war %d", spy.updateCalls)
	}
}

func TestModel_ReturnBook(t *testing.T) {
	model := NewModel(NewInMemoryRepository())
	model.AddBook(Book{ISBN: "111", Title: "Clean Code", Borrowed: true})

	book := model.ReturnBook("111")

	if book == nil {
		t.Fatal("ReturnBook gab nil zurück - erwartet das zurückgegebene Buch")
	}
	if book.Borrowed {
		t.Error("Buch sollte nach Rückgabe nicht mehr ausgeliehen sein")
	}
}

func TestModel_CheckBookAvailability(t *testing.T) {
	model := NewModel(NewInMemoryRepository())
	model.AddBook(Book{ISBN: "111", Title: "Clean Code"})

	available := model.CheckBookAvailability("111")

	if !available {
		t.Error("Buch sollte verfügbar sein")
	}
}

func TestModel_AvailableBooks(t *testing.T) {
	model := NewModel(NewInMemoryRepository())
	model.AddBook(Book{ISBN: "1", Borrowed: true})
	model.AddBook(Book{ISBN: "2", Borrowed: false})
	model.AddBook(Book{ISBN: "3", Borrowed: true})

	available := model.AvailableBooks()

	if len(available) != 1 {
		t.Fatalf("erwartet 1 verfügbares Buch, bekommen %d", len(available))
	}
	if available[0].ISBN != "2" {
		t.Errorf("falsches Buch: erwartete ISBN '2', bekommen '%s'", available[0].ISBN)
	}
}

func TestModel_BorrowedBooks(t *testing.T) {
	model := NewModel(NewInMemoryRepository())
	model.AddBook(Book{ISBN: "1", Borrowed: true})
	model.AddBook(Book{ISBN: "2", Borrowed: false})
	model.AddBook(Book{ISBN: "3", Borrowed: true})

	borrowed := model.BorrowedBooks()

	if len(borrowed) != 2 {
		t.Fatalf("erwartet 2 ausgeliehene Bücher, bekommen %d", len(borrowed))
	}
	if borrowed[0].ISBN != "1" || borrowed[1].ISBN != "3" {
		t.Errorf("falsche Bücher: erwartete '1' und '3', bekommen '%s' und '%s'", borrowed[0].ISBN, borrowed[1].ISBN)
	}
}

func TestModel_FindDuplicateISBN(t *testing.T) {
	model := NewModel(NewInMemoryRepository())
	model.AddBook(Book{ISBN: "111", Title: "Clean Code"})
	model.AddBook(Book{ISBN: "111", Title: "Clean Code 2"})
	model.AddBook(Book{ISBN: "222", Title: "The Pragmatic Programmer"})

	duplicate := model.FindDuplicateISBN("111")

	if !duplicate {
		t.Error("erwartet, dass eine doppelte ISBN gefunden wird")
	}
}

func TestDuplicateISBN(t *testing.T) {
	model := NewModel(NewInMemoryRepository())
	model.AddBook(Book{ISBN: "111", Title: "Clean Code"})
	model.AddBook(Book{ISBN: "111", Title: "Clean Code 2"})
	model.AddBook(Book{ISBN: "222", Title: "The Pragmatic Programmer"})

	duplicate := model.FindDuplicateISBN("111")

	if !duplicate {
		t.Errorf("erwartet, dass ein doppeltes ISBN gefunden wird")
	}
}
