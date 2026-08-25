package view

import (
	"testing"

	"mvc-booklibrary/model"
)

func TestView_PrintBookList(t *testing.T) {
	tests := []struct {
		name         string
		books        []model.Book
		wantContains []string
	}{
		{
			name: "mehrere Bücher",
			books: []model.Book{
				{ISBN: "111", Title: "Clean Code", Author: "Martin", PublishedYear: 2008, Borrowed: false},
				{ISBN: "222", Title: "The Pragmatic Programmer", Author: "Hunt", PublishedYear: 1999, Borrowed: true},
			},
			wantContains: []string{"111", "Clean Code", "BORROWED: false", "222", "The Pragmatic Programmer", "BORROWED: true"},
		},
		{
			name:         "einzelnes Buch",
			books:        []model.Book{{ISBN: "42", Title: "Der Kaktus", Author: "Blume", PublishedYear: 2020, Borrowed: false}},
			wantContains: []string{"42", "Der Kaktus", "BORROWED: false"},
		},
		{
			name:         "leere Liste",
			books:        []model.Book{},
			wantContains: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					t.Error("PrintBookList darf keine Panik auslösen")
				}
			}()

			PrintBookList(tt.books)
		})
	}
}
