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
