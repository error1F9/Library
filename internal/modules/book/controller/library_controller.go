package controller

import (
	"Library/internal/models"
	"Library/internal/modules/book/service"
	"Library/internal/responder"
	"errors"
	"github.com/ptflp/godecoder"
	"net/http"
)

type BookController struct {
	service service.BookServiceInterface
	godecoder.Decoder
	responder.Responder
}

type BookControllerInterface interface {
	AddBook(w http.ResponseWriter, r *http.Request)
	GetAllBooks(w http.ResponseWriter, r *http.Request)
}

func NewBookController(service service.BookServiceInterface, Decoder godecoder.Decoder, Responder responder.Responder) *BookController {
	return &BookController{service: service, Decoder: Decoder, Responder: Responder}
}

// AddBook
//
//	@Summary		Add a book
//	@Description	Adding a book
//	@Tags			Books
//	@Accept			json
//	@Produce		json
//	@Param			book	body		controller.AddBookRequest	true	"Book information with author id"
//	@Success		200		{object}	controller.AddBookResponse
//	@Failure		400		{string}	string	"error"
//	@Failure		404		{string}	string	"error"
//	@Failure		500		{string}	string	"error"
//	@Router			/books/add [post]
func (c *BookController) AddBook(w http.ResponseWriter, r *http.Request) {
	var bookDTO models.BookDTO
	err := c.Decode(r.Body, &bookDTO)
	if err != nil {
		c.ErrorBadRequest(w, err)
		return
	}

	if bookDTO.Author == nil {
		c.ErrorBadRequest(w, errors.New("Author ID is required"))
		return
	}

	book := &models.Book{
		Title:    bookDTO.Title,
		AuthorID: bookDTO.Author.ID,
	}

	out := c.service.AddBook(service.AddBookIn{AuthorID: book.AuthorID, BookTitle: book.Title})
	if out.Error != nil {
		c.OutputJSON(w, AddBookResponse{
			Success: false,
			Error:   out.Error,
		})
		return
	}

	bookDTO = models.BookDTO{
		ID:    out.Book.ID,
		Title: out.Book.Title,
		Author: &models.AuthorDTO{
			ID:        out.Author.ID,
			Firstname: out.Author.Firstname,
			Lastname:  out.Author.Lastname,
		},
	}

	c.OutputJSON(w, AddBookResponse{
		Success: true,
		Data: &DataBook{
			Message: "Book added successfully",
			Book:    bookDTO,
		},
	})
}

// GetAllBooks
//
//	@Summary		Get All Books
//	@Description	Getting all Books
//	@Tags			Books
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	controller.GetAllBooksResponse
//	@Failure		400	{string}	string	"error"
//	@Failure		404	{string}	string	"error"
//	@Failure		500	{string}	string	"error"
//	@Router			/books/all [get]
func (c *BookController) GetAllBooks(w http.ResponseWriter, r *http.Request) {
	out := c.service.GetAllBooks()
	if out.Error != nil {
		c.OutputJSON(w, GetAllBooksResponse{
			Success: false,
			Error:   out.Error.Error(),
		})
		return
	}
	bookDTOs := make([]models.BookDTO, len(out.Books))
	for i := 0; i < len(out.Books); i++ {
		bookDTOs[i].ID = out.Books[i].ID
		bookDTOs[i].Title = out.Books[i].Title
		bookDTOs[i].Author = &models.AuthorDTO{}
		bookDTOs[i].Author.ID = out.Books[i].Author.ID
		bookDTOs[i].Author.Firstname = out.Books[i].Author.Firstname
		bookDTOs[i].Author.Lastname = out.Books[i].Author.Lastname
		bookDTOs[i].Author.Rating = out.Books[i].Author.Rating
	}

	c.OutputJSON(w, GetAllBooksResponse{
		Success: true,
		Data: &DataBooks{
			Message: "All books from Library",
			Books:   bookDTOs,
		},
	})
}
