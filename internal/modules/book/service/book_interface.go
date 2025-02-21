package service

import (
	"Library/internal/models"
)

type BookServiceInterface interface {
	AddBook(in AddBookIn) AddBookOut
	GetAllBooks() GetAllBooksOut
	Empty() (bool, error)
}

type AddBookIn struct {
	AuthorID  uint   `json:"author_id"`
	BookTitle string `json:"book_title"`
}

type AddBookOut struct {
	Book   *models.Book   `json:"book"`
	Author *models.Author `json:"author"`
	Error  error          `json:"error,omitempty"`
}

type GetAllBooksOut struct {
	Books []*models.Book `json:"book"`
	Error error          `json:"error,omitempty"`
}
