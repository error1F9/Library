package service

import (
	"Library/internal/models"
	"context"
)

type UserServiceInterface interface {
	CreateUser(in CreateUserIn) CreateUserOut
	RentBookByUser(in RentBookByUserIn) error
	GetAllUsers() GetAllUsersOut
	ReturnBook(in ReturnBookIn) error
	Login(in LoginIn) LoginOut
	Logout(ctx context.Context) error
	Empty() (bool, error)
}

type RentBookByUserIn struct {
	BookID uint `json:"bookID"`
	UserID uint `json:"userID"`
}

type CreateUserIn struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type CreateUserOut struct {
	UserID uint  `json:"userID,omitempty"`
	Error  error `json:"error,omitempty"`
}

type GetAllUsersOut struct {
	Users []*models.User `json:"users,omitempty"`
	Error error          `json:"error,omitempty"`
}

type ReturnBookIn struct {
	BookID uint  `json:"bookID"`
	UserID uint  `json:"userID"`
	Error  error `json:"error,omitempty"`
}

type LoginOut struct {
	AccessToken string `json:"access_token"`
	Err         error  `json:"error,omitempty"`
}

type LoginIn struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
