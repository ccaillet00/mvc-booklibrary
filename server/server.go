package server

import (
	"encoding/json"
	"net/http"

	"mvc-booklibrary/model"
)

// Server exposes the book library as an HTTP API.
type Server struct {
	lib *model.Model
}

// NewServer creates a new HTTP server with the given model.
func NewServer(lib *model.Model) *Server {
	return &Server{lib: lib}
}

// Routes returns the HTTP mux with all registered routes.
func (s *Server) Routes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /books", s.listBooks)
	mux.HandleFunc("GET /books/{isbn}", s.getBook)
	mux.HandleFunc("POST /books", s.addBook)
	mux.HandleFunc("DELETE /books/{isbn}", s.removeBook)
	mux.HandleFunc("POST /books/{isbn}/lend", s.lendBook)
	mux.HandleFunc("POST /books/{isbn}/return", s.returnBook)
	return mux
}

// listBooks returns all books as JSON.
func (s *Server) listBooks(w http.ResponseWriter, r *http.Request) {
	books := s.lib.FindAllBooks()
	writeJSON(w, http.StatusOK, books)
}

// getBook returns a single book by ISBN.
func (s *Server) getBook(w http.ResponseWriter, r *http.Request) {
	isbn := r.PathValue("isbn")
	book := s.lib.FindAllBooks()
	for _, b := range book {
		if b.ISBN == isbn {
			writeJSON(w, http.StatusOK, b)
			return
		}
	}
	writeJSON(w, http.StatusNotFound, map[string]string{"error": "book not found"})
}

// addBook adds a new book from a JSON body.
func (s *Server) addBook(w http.ResponseWriter, r *http.Request) {
	var book model.Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid JSON body"})
		return
	}
	s.lib.AddBook(book)
	writeJSON(w, http.StatusCreated, book)
}

// removeBook removes a book by ISBN.
func (s *Server) removeBook(w http.ResponseWriter, r *http.Request) {
	isbn := r.PathValue("isbn")
	if !s.lib.RemoveBook(isbn) {
		writeJSON(w, http.StatusNotFound, map[string]string{"error": "book not found"})
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

// lendBook marks a book as borrowed.
func (s *Server) lendBook(w http.ResponseWriter, r *http.Request) {
	isbn := r.PathValue("isbn")
	book := s.lib.LendBook(isbn)
	if book == nil {
		writeJSON(w, http.StatusNotFound, map[string]string{"error": "book not found"})
		return
	}
	writeJSON(w, http.StatusOK, book)
}

// returnBook marks a book as returned.
func (s *Server) returnBook(w http.ResponseWriter, r *http.Request) {
	isbn := r.PathValue("isbn")
	book := s.lib.ReturnBook(isbn)
	if book == nil {
		writeJSON(w, http.StatusNotFound, map[string]string{"error": "book not found or not borrowed"})
		return
	}
	writeJSON(w, http.StatusOK, book)
}

// writeJSON encodes v as JSON and writes it with the given status code.
func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)
}
