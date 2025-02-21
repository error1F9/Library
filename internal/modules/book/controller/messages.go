package controller

import (
	"Library/internal/models"
)

type AddBookRequest struct {
	Title    string `json:"title"`
	AuthorID string `json:"author_id"`
}

type AddBookResponse struct {
	Success bool      `json:"success"`
	Error   error     `json:"error,omitempty"`
	Data    *DataBook `json:"data,omitempty"`
}

type DataBook struct {
	Message string         `json:"message,omitempty"`
	Book    models.BookDTO `json:"book,omitempty"`
}

type GetAllBooksResponse struct {
	Success bool       `json:"success"`
	Error   string     `json:"error,omitempty"`
	Data    *DataBooks `json:"data,omitempty"`
}

type DataBooks struct {
	Message string           `json:"message,omitempty"`
	Books   []models.BookDTO `json:"books,omitempty"`
}
