package main

import (
	"encoding/json"
	"net/http"
)

type Server struct {
	lib Model
}

func NewServer(lib Model) *Server {
	return &Server{lib: lib}
}

func (s *Server) Routes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /books", s.listBooks)
	return mux
}

func (s *Server) listBooks(w http.ResponseWriter, r *http.Request) {
	books := s.lib.FindAllBooks()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(books)
}
