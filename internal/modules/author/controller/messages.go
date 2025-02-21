package controller

import (
	"Library/internal/models"
)

type AddAuthorRequest struct {
	Firstname string `json:"firstname,example:Alexandr"`
	Lastname  string `json:"lastname,example:Pushkin"`
}

type AddAuthorResponse struct {
	Success bool        `json:"success"`
	Error   error       `json:"error,omitempty"`
	Data    *DataAuthor `json:"data,omitempty"`
}

type DataAuthor struct {
	Message string           `json:"message,omitempty"`
	Author  models.AuthorDTO `json:"author,omitempty"`
}

type GetAllAuthorsResponse struct {
	Success bool         `json:"success"`
	Error   string       `json:"error,omitempty"`
	Data    *DataAuthors `json:"data,omitempty"`
}

type DataAuthors struct {
	Message string             `json:"message,omitempty"`
	Authors []models.AuthorDTO `json:"authors,omitempty"`
}

type GetTopAuthorsResponse struct {
	Success bool         `json:"success"`
	Error   string       `json:"error,omitempty"`
	Data    *DataAuthors `json:"data,omitempty"`
}
