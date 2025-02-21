package controller

import (
	"Library/internal/models"
	"Library/internal/modules/author/service"
	"Library/internal/responder"
	"fmt"
	"github.com/ptflp/godecoder"
	"net/http"
	"strconv"
)

type AuthorController struct {
	service service.AuthorServiceInterface
	godecoder.Decoder
	responder.Responder
}

type AuthorControllerInterface interface {
	AddAuthor(w http.ResponseWriter, r *http.Request)
	GetAllAuthors(w http.ResponseWriter, r *http.Request)
	GetTopAuthors(w http.ResponseWriter, r *http.Request)
}

func NewAuthorController(service service.AuthorServiceInterface, Decoder godecoder.Decoder, Responder responder.Responder) *AuthorController {
	return &AuthorController{service: service, Decoder: Decoder, Responder: Responder}
}

// AddAuthor
//
//	@Summary		Add an author
//	@Description	Adding an author
//	@Tags			Authors
//	@Accept			json
//	@Produce		json
//	@Param			author	body		controller.AddAuthorRequest	true	"Author information"
//	@Success		200		{object}	controller.AddAuthorResponse
//	@Failure		400		{string}	string	"error"
//	@Failure		404		{string}	string	"error"
//	@Failure		500		{string}	string	"error"
//	@Router			/authors/add [post]
func (c *AuthorController) AddAuthor(w http.ResponseWriter, r *http.Request) {
	var authorDTO models.AuthorDTO
	err := c.Decode(r.Body, &authorDTO)
	if err != nil {
		c.ErrorBadRequest(w, err)
		return
	}

	author := &models.Author{
		Firstname: authorDTO.Firstname,
		Lastname:  authorDTO.Lastname,
	}

	out := c.service.AddAuthor(service.AddAuthorIn{Firstname: author.Firstname, Lastname: author.Lastname})
	if out.Error != nil {
		c.OutputJSON(w, AddAuthorResponse{
			Success: false,
			Error:   out.Error,
		})
		return
	}

	authorDTO = models.AuthorDTO{
		ID:        out.Author.ID,
		Firstname: out.Author.Firstname,
		Lastname:  out.Author.Lastname,
	}

	c.OutputJSON(w, AddAuthorResponse{
		Success: true,
		Data: &DataAuthor{
			Message: "Author added successfully",
			Author:  authorDTO,
		},
	})
}

// GetTopAuthors
//
//	@Summary		Get top authors
//	@Description	Getting a top of author with your limit
//	@Tags			Authors
//	@Accept			json
//	@Produce		json
//	@Param			limit	query		int	false	"This is how many authors need to get"
//	@Success		200		{object}	controller.GetTopAuthorsResponse
//	@Failure		400		{string}	string	"error"
//	@Failure		404		{string}	string	"error"
//	@Failure		500		{string}	string	"error"
//	@Router			/authors/top [get]
func (c *AuthorController) GetTopAuthors(w http.ResponseWriter, r *http.Request) {
	limitString := r.URL.Query().Get("limit")
	limit, err := strconv.Atoi(limitString)
	if err != nil {
		c.ErrorBadRequest(w, err)
	}

	out := c.service.GetTopAuthors(limit)
	if out.Error != nil {
		c.OutputJSON(w, GetTopAuthorsResponse{
			Success: false,
			Error:   out.Error.Error(),
		})
		return
	}

	authorDTOs := make([]models.AuthorDTO, len(out.Authors))
	for i := 0; i < len(out.Authors); i++ {
		authorDTOs[i].ID = out.Authors[i].ID
		authorDTOs[i].Firstname = out.Authors[i].Firstname
		authorDTOs[i].Lastname = out.Authors[i].Lastname
		authorDTOs[i].Rating = out.Authors[i].Rating

	}

	c.OutputJSON(w, GetAllAuthorsResponse{
		Success: true,
		Data: &DataAuthors{
			Message: fmt.Sprintf("Top %d authors from Library", limit),
			Authors: authorDTOs,
		},
	})
}

// GetAllAuthors
//
//	@Summary		Get All authors
//	@Description	Getting all authors
//	@Tags			Authors
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	controller.GetAllAuthorsResponse
//	@Failure		400	{string}	string	"error"
//	@Failure		404	{string}	string	"error"
//	@Failure		500	{string}	string	"error"
//	@Router			/authors/all [get]
func (c *AuthorController) GetAllAuthors(w http.ResponseWriter, r *http.Request) {
	out := c.service.GetAllAuthors()
	if out.Error != nil {
		c.OutputJSON(w, GetAllAuthorsResponse{
			Success: false,
			Error:   out.Error.Error(),
		})
		return
	}
	authorDTOs := make([]models.AuthorDTO, len(out.Authors))
	for i := 0; i < len(out.Authors); i++ {
		authorDTOs[i].ID = out.Authors[i].ID
		authorDTOs[i].Firstname = out.Authors[i].Firstname
		authorDTOs[i].Lastname = out.Authors[i].Lastname
		authorDTOs[i].Rating = out.Authors[i].Rating

		authorDTOs[i].Books = make([]*models.BookDTO, len(out.Authors[i].Books))
		for j := 0; j < len(out.Authors[i].Books); j++ {
			authorDTOs[i].Books[j] = &models.BookDTO{}
			authorDTOs[i].Books[j].ID = out.Authors[i].Books[j].ID
			authorDTOs[i].Books[j].Title = out.Authors[i].Books[j].Title
		}
	}

	c.OutputJSON(w, GetAllAuthorsResponse{
		Success: true,
		Data: &DataAuthors{
			Message: "All authors from Library",
			Authors: authorDTOs,
		},
	})
}
