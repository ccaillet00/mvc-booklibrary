package gin_server

import (
	"mvc-booklibrary/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GinServer exposes the book library as an HTTP API using the Gin framework.
type GinServer struct {
	lib *model.Model
}

// NewGinServer creates a new HTTP server with the given model.
func NewGinServer(lib *model.Model) *GinServer {
	return &GinServer{lib: lib}
}

// Routes returns the Gin engine with all registered routes.
func (s *GinServer) Routes() *gin.Engine {
	r := gin.Default()
	r.Use(gin.Logger(), gin.Recovery())

	books := r.Group("/books")

	books.GET("", s.listBooks)
	books.GET("/:isbn", s.getBook)
	books.POST("", s.addBook)
	books.DELETE("/:isbn", s.removeBook)
	books.POST("/:isbn/lend", s.lendBook)
	books.POST("/:isbn/return", s.returnBook)
	return r
}

func (s *GinServer) listBooks(context *gin.Context) {
	context.JSON(http.StatusOK, s.lib.FindAllBooks())
}

// getBook returns a single book by ISBN.
func (s *GinServer) getBook(context *gin.Context) {
	isbn := context.Param("isbn")
	book := s.lib.FindBook(isbn)
	if book == nil {
		fehler(context, http.StatusNotFound, "book not found")
		return
	}
	context.JSON(http.StatusOK, book)
}

// addBook adds a new book from a JSON body.
func (s *GinServer) addBook(context *gin.Context) {
	var book model.Book
	if err := context.ShouldBindJSON(&book); err != nil {
		fehler(context, http.StatusBadRequest, "invalid JSON body")
		return
	}
	if s.lib.FindBook(book.ISBN) != nil {
		fehler(context, http.StatusConflict, "book with this ISBN already exists")
		return
	}

	if book.ISBN == "" || book.Title == "" || book.Author == "" || book.PublishedYear == 0 {
		fehler(context, http.StatusBadRequest, "missing required fields")
		return
	}
	s.lib.AddBook(book)
	context.JSON(http.StatusCreated, book)
}

// removeBook removes a book by ISBN.
func (s *GinServer) removeBook(context *gin.Context) {
	isbn := context.Param("isbn")
	if !s.lib.RemoveBook(isbn) {
		fehler(context, http.StatusNotFound, "book not found")
		return
	}
	context.Writer.WriteHeader(http.StatusNoContent)
}

// lendBook marks a book as borrowed.
func (s *GinServer) lendBook(context *gin.Context) {
	isbn := context.Param("isbn")
	book := s.lib.LendBook(isbn)
	if book == nil {
		fehler(context, http.StatusNotFound, "book not found")
		return
	}
	context.JSON(http.StatusOK, book)
}

// returnBook marks a book as returned.
func (s *GinServer) returnBook(context *gin.Context) {
	isbn := context.Param("isbn")
	book := s.lib.ReturnBook(isbn)
	if book == nil {
		fehler(context, http.StatusNotFound, "book not found or not borrowed")
		return
	}
	context.JSON(http.StatusOK, book)
}

func fehler(c *gin.Context, status int, message string) {
	c.JSON(status, gin.H{"error": message})
}
