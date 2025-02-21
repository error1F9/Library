package controller

import (
	"Library/internal/models"
	"Library/internal/modules/user/service"
	"Library/internal/responder"
	"fmt"
	"github.com/ptflp/godecoder"
	"log"
	"net/http"
	"strconv"
)

type UserController struct {
	service service.UserServiceInterface
	godecoder.Decoder
	responder.Responder
}

type UserControllerInterface interface {
	CreateUser(w http.ResponseWriter, r *http.Request)
	GetAllUsers(w http.ResponseWriter, r *http.Request)
	RentBookByUser(w http.ResponseWriter, r *http.Request)
	ReturnBook(w http.ResponseWriter, r *http.Request)
	Login(w http.ResponseWriter, r *http.Request)
	Logout(w http.ResponseWriter, r *http.Request)
}

func NewUserController(service service.UserServiceInterface, decoder godecoder.Decoder, responder responder.Responder) *UserController {
	return &UserController{service: service, Responder: responder, Decoder: decoder}
}

// CreateUser godoc
//
//	@Summary		Create a new user
//	@Description	Create a new user with the provided information
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Param			user	body		controller.CreateUserRequest	true	"User information"
//	@Success		200		{object}	controller.CreateUserResponse
//	@Failure		400		{object}	controller.ErrorResponse
//	@Router			/user/create [post]
func (c *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	var userDTO models.UserDTO
	if err := c.Decode(r.Body, &userDTO); err != nil {
		c.ErrorBadRequest(w, err)
		return
	}

	out := c.service.CreateUser(service.CreateUserIn{
		Firstname: userDTO.Firstname,
		Lastname:  userDTO.Lastname,
		Email:     userDTO.Email,
		Password:  userDTO.Password,
	})

	if out.Error != nil {
		c.OutputJSON(w, CreateUserResponse{
			Success: false,
			Error:   out.Error,
		})
	}

	c.OutputJSON(w, CreateUserResponse{
		Success: true,
		Message: fmt.Sprintf("User created with id: %d", out.UserID),
	})
}

// RentBookByUser godoc
//
//	@Summary		Rent a book by user
//	@Description	Rent a book by a user with the provided user ID and book ID
//	@Tags			User
//	@Security		ApiKeyAuth
//	@Accept			json
//	@Produce		json
//	@Param			user_id	query		int	true	"User ID"
//	@Param			book_id	query		int	true	"Book ID"
//	@Success		200		{object}	controller.RentBookByUserResponse
//	@Failure		400		{object}	controller.ErrorResponse
//	@Router			/user/rent [post]
func (c *UserController) RentBookByUser(w http.ResponseWriter, r *http.Request) {
	userIDfromQuery := r.URL.Query().Get("user_id")
	bookIDfromQuery := r.URL.Query().Get("book_id")

	userID, err := strconv.Atoi(userIDfromQuery)
	if err != nil {
		c.ErrorBadRequest(w, err)
		return
	}

	bookID, err := strconv.Atoi(bookIDfromQuery)
	if err != nil {
		c.ErrorBadRequest(w, err)
		return
	}

	if err = c.service.RentBookByUser(service.RentBookByUserIn{
		BookID: uint(bookID),
		UserID: uint(userID),
	}); err != nil {
		c.OutputJSON(w, RentBookByUserResponse{
			Success: false,
			Error:   err.Error(),
		})
		return
	}

	c.OutputJSON(w, RentBookByUserResponse{
		Success: true,
		Message: fmt.Sprintf("User %d Rented book with id: %d", userID, bookID),
	})
}

// ReturnBook godoc
//
//	@Summary		Return a book by user
//	@Description	Return a book by a user with the provided user ID and book ID
//	@Tags			User
//	@Security		ApiKeyAuth
//	@Accept			json
//	@Produce		json
//	@Param			user_id	query		int	true	"User ID"
//	@Param			book_id	query		int	true	"Book ID"
//	@Success		200		{object}	controller.RentBookByUserResponse
//	@Failure		400		{object}	controller.ErrorResponse
//	@Router			/user/return [post]
func (c *UserController) ReturnBook(w http.ResponseWriter, r *http.Request) {
	userIDfromQuery := r.URL.Query().Get("user_id")
	bookIDfromQuery := r.URL.Query().Get("book_id")

	userID, err := strconv.Atoi(userIDfromQuery)
	if err != nil {
		c.ErrorBadRequest(w, err)
		return
	}

	bookID, err := strconv.Atoi(bookIDfromQuery)
	if err != nil {
		c.ErrorBadRequest(w, err)
		return
	}

	if err = c.service.ReturnBook(service.ReturnBookIn{
		BookID: uint(bookID),
		UserID: uint(userID),
	}); err != nil {
		c.OutputJSON(w, RentBookByUserResponse{
			Success: false,
			Error:   err.Error(),
		})
		return
	}

	c.OutputJSON(w, RentBookByUserResponse{
		Success: true,
		Message: fmt.Sprintf("User %d Returned book with id: %d", userID, bookID),
	})
}

// GetAllUsers
//
//	@Summary		Get All Users
//	@Description	Getting all Users
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	controller.GetAllUsersResponse
//	@Failure		400	{string}	string	"error"
//	@Failure		404	{string}	string	"error"
//	@Failure		500	{string}	string	"error"
//	@Router			/user/all [get]
func (c *UserController) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	out := c.service.GetAllUsers()
	if out.Error != nil {
		c.OutputJSON(w, GetAllUsersResponse{
			Success: false,
			Error:   out.Error.Error(),
		})
		return
	}

	userDTOs := make([]models.UserDTO, len(out.Users))
	for i, user := range out.Users {
		userDTOs[i].ID = user.ID
		userDTOs[i].Firstname = user.Firstname
		userDTOs[i].Lastname = user.Lastname
		userDTOs[i].Email = user.Email
		userDTOs[i].Password = user.Password
	}

	c.OutputJSON(w, GetAllUsersResponse{
		Success: true,
		Data: &DataUsers{
			Message: "All users",
			Users:   userDTOs,
		},
	})
}

// Login
//
//	@Summary		User login
//	@Description	Authenticate a user by email and password
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Param			email		query		string	true	"Email"
//	@Param			password	query		string	true	"Password"
//	@Success		200			{string}	string	"Access token"
//	@Router			/user/login [post]
func (c *UserController) Login(w http.ResponseWriter, r *http.Request) {
	email := r.URL.Query().Get("email")
	password := r.URL.Query().Get("password")
	log.Println(email)

	auth := c.service.Login(service.LoginIn{Email: email, Password: password})
	if auth.Err != nil {
		c.ErrorUnauthorized(w, auth.Err)
		return
	}
	accessToken := auth.AccessToken

	c.OutputJSON(w, fmt.Sprintf("Bearer %s", accessToken))

}

// Logout
//
//	@Summary		User logout
//	@Description	Terminate the user session
//	@Tags			User
//	@Security		ApiKeyAuth
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	controller.LogoutResponse	"Logout success"
//	@Router			/user/logout [post]
func (c *UserController) Logout(w http.ResponseWriter, r *http.Request) {
	err := c.service.Logout(r.Context())
	if err != nil {
		c.OutputJSON(w, LogoutResponse{
			Success: false,
			Error:   err.Error(),
			Data:    "Logout error: " + err.Error(),
		})
		return
	}
	c.OutputJSON(w, LogoutResponse{
		Success: true,
		Data:    "Logout success",
	})
}
