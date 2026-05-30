package main

import "testing"

func TestCreateBook(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    Book
		wantErr bool
	}{
		{"mit Leerzeichen", "111, Clean Code, Martin, 2008",
			Book{"111", "Clean Code", "Martin", 2008, false}, false},
		{"ohne Leerzeichen", "111,Clean Code,Martin,2008",
			Book{"111", "Clean Code", "Martin", 2008, false}, false},
		{"mit und ohne Leerzeichen", "111,Clean Code, Martin, 2008",
			Book{"111", "Clean Code", "Martin", 2008, false}, false},
		{"fehlende Felder", "111, Clean Code", Book{}, true},
		{"Jahr keine Zahl", "1, A, B, zweitausend, true", Book{}, true},
	}
	c := NewController(NewModel(NewInMemoryRepository()))
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := c.createBook(tc.input)
			if (err != nil) != tc.wantErr {
				t.Errorf("createBook(%q): err = %v, wantErr = %v", tc.input, err, tc.wantErr)
			}
			if got != tc.want {
				t.Errorf("createBook(%q) = %v, want %v", tc.input, got, tc.want)
			}
		})
	}
}
