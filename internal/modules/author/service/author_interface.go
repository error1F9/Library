package service

import (
	"Library/internal/models"
)

type AuthorServiceInterface interface {
	AddAuthor(in AddAuthorIn) AddAuthorOut
	GetAllAuthors() GetAllAuthorsOut
	GetTopAuthors(limit int) GetAllAuthorsOut
	Empty() (bool, error)
}

type AddAuthorIn struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

type AddAuthorOut struct {
	Author *models.Author `json:"author"`
	Error  error          `json:"error,omitempty"`
}

type GetAllAuthorsOut struct {
	Authors []*models.Author `json:"authors"`
	Error   error            `json:"error,omitempty"`
}

type GetTopAuthorsOut struct {
	Authors []*models.Author `json:"authors"`
	Error   error            `json:"error,omitempty"`
}
