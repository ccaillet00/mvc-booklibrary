package controller

import (
	"testing"

	"mvc-booklibrary/model"
	"mvc-booklibrary/repository"
)

// ============================================================================
// Controller-Tests (Unit-Tests)
// Testen die Eingabeverarbeitung des Controllers isoliert.
// ============================================================================

func TestController_CreateBook(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    model.Book
		wantErr bool
	}{
		{
			name:    "mit Leerzeichen",
			input:   "111, Clean Code, Martin, 2008",
			want:    model.Book{ISBN: "111", Title: "Clean Code", Author: "Martin", PublishedYear: 2008, Borrowed: false},
			wantErr: false,
		},
		{
			name:    "ohne Leerzeichen",
			input:   "111,Clean Code,Martin,2008",
			want:    model.Book{ISBN: "111", Title: "Clean Code", Author: "Martin", PublishedYear: 2008, Borrowed: false},
			wantErr: false,
		},
		{
			name:    "gemischte Leerzeichen",
			input:   "111,Clean Code, Martin, 2008",
			want:    model.Book{ISBN: "111", Title: "Clean Code", Author: "Martin", PublishedYear: 2008, Borrowed: false},
			wantErr: false,
		},
		{
			name:    "fehlende Felder",
			input:   "111, Clean Code",
			want:    model.Book{},
			wantErr: true,
		},
		{
			name:    "Jahr keine Zahl",
			input:   "1, A, B, zweitausend, true",
			want:    model.Book{},
			wantErr: true,
		},
	}

	ctrl := NewController(model.NewModel(repository.NewInMemoryRepository()))

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ctrl.CreateBook(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateBook(%q): err = %v, wantErr = %v", tt.input, err, tt.wantErr)
			}
			if got != tt.want {
				t.Errorf("CreateBook(%q) = %+v, want %+v", tt.input, got, tt.want)
			}
		})
	}
}
