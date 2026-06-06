package main

import "testing"

type SpyRepository struct {
	Repository
	updateCalls int
}

func (s *SpyRepository) Update(book Book) bool {
	s.updateCalls++ //mitzählen
	return s.Repository.Update(book)
}

func TestAddBook(t *testing.T) {

	// Arrange - Ausgangslage aufbauen
	repo := NewInMemoryRepository()
	model := NewModel(repo)

	// Act - die zu testende Aktion
	model.AddBook(Book{ISBN: "111", Title: "Clean Code"})

	// Assert - Ergebnis behaupten
	books := model.FindAllBooks()
	if len(books) != 1 {
		t.Errorf("erwartet 1 Buch, bekommen %d", len(books))
	}
}

func TestLendBook(t *testing.T) {
	spy := &SpyRepository{Repository: NewInMemoryRepository()}
	model := NewModel(spy)
	model.AddBook(Book{ISBN: "111", Title: "Clean Code"})

	model.LendBook("111")

	if spy.updateCalls != 1 {
		t.Errorf("LendBook sollte Update 1x aufrufen, war %d", spy.updateCalls)
	}
}

func TestReturnBook(t *testing.T) {
	model := NewModel(NewInMemoryRepository())
	model.AddBook(Book{ISBN: "111", Title: "Clean Code", Borrowed: true})

	book := model.ReturnBook("111")

	if book == nil {
		t.Fatal("ReturnBook gab nil zurück - erwartet das ausgeliehene Buch")
	}
	if book.Borrowed {
		t.Errorf("Buch sollte zurückgegeben markiert sein")
	}
}

func TestCheckBookAvailability(t *testing.T) {
	model := NewModel(NewInMemoryRepository())
	model.AddBook(Book{ISBN: "111", Title: "Clean Code"})

	book := model.CheckBookAvailability("111")

	if !book {
		t.Errorf("Buch sollte verfügbar markiert sein")
	}
}

func TestAvailableBook(t *testing.T) {
	model := NewModel(NewInMemoryRepository())
	model.AddBook(Book{ISBN: "1", Borrowed: true})
	model.AddBook(Book{ISBN: "2", Borrowed: false})
	model.AddBook(Book{ISBN: "3", Borrowed: true})

	available := model.AvailableBooks()

	if len(available) != 1 {
		t.Fatalf("erwartet 1 verfügbares Buch, bekommen %d", len(available))
	}

	if available[0].ISBN != "2" {
		t.Errorf("falsche Buch zurückgeben: %s", available[0].ISBN)
	}
}

func TestBorrowedBook(t *testing.T) {
	model := NewModel(NewInMemoryRepository())
	model.AddBook(Book{ISBN: "1", Borrowed: true})
	model.AddBook(Book{ISBN: "2", Borrowed: false})
	model.AddBook(Book{ISBN: "3", Borrowed: true})

	borrowed := model.BorrowedBooks()

	if len(borrowed) != 2 {
		t.Fatalf("erwartet 2 ausgeliehene Bücher, bekommen %d", len(borrowed))
	}

	if borrowed[0].ISBN != "1" || borrowed[1].ISBN != "3" {
		t.Errorf("falsche Bücher zurückgeben: %s, %s", borrowed[0].ISBN, borrowed[1].ISBN)
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


