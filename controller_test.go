package main

import (
	"testing"
)

// ============================================================================
// Controller-Tests (Unit-Tests)
// Testen die Eingabeverarbeitung des Controllers isoliert.
// ============================================================================

func TestController_CreateBook(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    Book
		wantErr bool
	}{
		{
			name:    "mit Leerzeichen",
			input:   "111, Clean Code, Martin, 2008",
			want:    Book{ISBN: "111", Title: "Clean Code", Author: "Martin", PublishedYear: 2008, Borrowed: false},
			wantErr: false,
		},
		{
			name:    "ohne Leerzeichen",
			input:   "111,Clean Code,Martin,2008",
			want:    Book{ISBN: "111", Title: "Clean Code", Author: "Martin", PublishedYear: 2008, Borrowed: false},
			wantErr: false,
		},
		{
			name:    "gemischte Leerzeichen",
			input:   "111,Clean Code, Martin, 2008",
			want:    Book{ISBN: "111", Title: "Clean Code", Author: "Martin", PublishedYear: 2008, Borrowed: false},
			wantErr: false,
		},
		{
			name:    "fehlende Felder",
			input:   "111, Clean Code",
			want:    Book{},
			wantErr: true,
		},
		{
			name:    "Jahr keine Zahl",
			input:   "1, A, B, zweitausend, true",
			want:    Book{},
			wantErr: true,
		},
	}

	controller := NewController(NewModel(NewInMemoryRepository()))

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := controller.createBook(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("createBook(%q): err = %v, wantErr = %v", tt.input, err, tt.wantErr)
			}
			if got != tt.want {
				t.Errorf("createBook(%q) = %+v, want %+v", tt.input, got, tt.want)
			}
		})
	}
}
