package main

import (
	"testing"
)

func TestIntegration_AddBookAndCheckAvailability(t *testing.T) {
	repo := NewInMemoryRepository()
	model := NewModel(repo)
	controller := NewController(model)

	input := "123, Clean Code, Robert C. Martin, 2008"
	book, err := controller.createBook(input)
	if err != nil {
		t.Fatalf("createBook fehlgeschlagen: %v", err)
	}
	model.AddBook(book)

	if !model.CheckBookAvailability("123") {
		t.Error("Buch sollte verfügbar sein")
	}

	allBooks := model.FindAllBooks()
	if len(allBooks) != 1 {
		t.Errorf("erwartet 1 Buch, bekommen %d", len(allBooks))
	}
	if allBooks[0].ISBN != "123" || allBooks[0].Title != "Clean Code" {
		t.Errorf("falsches Buch: %+v", allBooks[0])
	}
}

func TestIntegration_LendAndReturnBook(t *testing.T) {
	repo := NewInMemoryRepository()
	model := NewModel(repo)

	model.AddBook(Book{ISBN: "456", Title: "Der C++ Programmierer", Author: "Stroustrup", PublishedYear: 2013, Borrowed: false})

	if !model.CheckBookAvailability("456") {
		t.Error("Buch sollte initial verfügbar sein")
	}

	borrowed := model.LendBook("456")
	if borrowed == nil {
		t.Fatal("LendBook gab nil zurück")
	}
	if !borrowed.Borrowed {
		t.Error("Buch sollte nach Ausleihe als 'Borrowed' markiert sein")
	}

	if model.CheckBookAvailability("456") {
		t.Error("Buch sollte nach Ausleihe nicht mehr verfügbar sein")
	}

	returned := model.ReturnBook("456")
	if returned == nil {
		t.Fatal("ReturnBook gab nil zurück")
	}
	if returned.Borrowed {
		t.Error("Buch sollte nach Rückgabe wieder verfügbar sein")
	}

	if !model.CheckBookAvailability("456") {
		t.Error("Buch sollte nach Rückgabe wieder verfügbar sein")
	}
}

func TestIntegration_RemoveBook(t *testing.T) {
	repo := NewInMemoryRepository()
	model := NewModel(repo)

	model.AddBook(Book{ISBN: "789", Title: "Design Patterns", Author: "Gang of Four", PublishedYear: 1994, Borrowed: false})

	if len(model.FindAllBooks()) != 1 {
		t.Error("Buch sollte im Repository existieren")
	}

	removed := model.RemoveBook("789")
	if !removed {
		t.Error("RemoveBook sollte true zurückgeben")
	}
	if len(model.FindAllBooks()) != 0 {
		t.Errorf("Repository sollte leer sein, aber enthält %d Bücher", len(model.FindAllBooks()))
	}

	removedAgain := model.RemoveBook("789")
	if removedAgain {
		t.Error("RemoveBook sollte false zurückgeben, wenn Buch nicht existiert")
	}
}

func TestIntegration_AvailableAndBorrowedBooks(t *testing.T) {
	repo := NewInMemoryRepository()
	model := NewModel(repo)

	model.AddBook(Book{ISBN: "1", Title: "Buch A", Borrowed: false})
	model.AddBook(Book{ISBN: "2", Title: "Buch B", Borrowed: true})
	model.AddBook(Book{ISBN: "3", Title: "Buch C", Borrowed: false})

	available := model.AvailableBooks()
	if len(available) != 2 {
		t.Fatalf("erwartet 2 verfügbare Bücher, bekommen %d", len(available))
	}

	borrowed := model.BorrowedBooks()
	if len(borrowed) != 1 {
		t.Fatalf("erwartet 1 ausgeliehenes Buch, bekommen %d", len(borrowed))
	}
	if borrowed[0].ISBN != "2" {
		t.Errorf("erwartet ISBN '2' als ausgeliehen, bekommen '%s'", borrowed[0].ISBN)
	}
}
