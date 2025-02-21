package controller

import "Library/internal/models"

type CreateUserRequest struct {
	Firstname string `json:"firstname,example:Peter"`
	Lastname  string `json:"lastname,example:Jackson"`
	Email     string `json:"email,example:123@example.com"`
	Password  string `json:"password,example:qwerty12"`
}

type CreateUserResponse struct {
	Success bool   `json:"success"`
	Error   error  `json:"error,omitempty"`
	Message string `json:"message,omitempty"`
}

type GetAllUsersResponse struct {
	Success bool       `json:"success"`
	Error   string     `json:"error,omitempty"`
	Data    *DataUsers `json:"data,omitempty"`
}

type DataUsers struct {
	Message string           `json:"message,omitempty"`
	Users   []models.UserDTO `json:"users,omitempty"`
}

type RentBookByUserResponse struct {
	Success bool   `json:"success"`
	Error   string `json:"error,omitempty"`
	Message string `json:"message,omitempty"`
}

type ErrorResponse struct {
	Success bool   `json:"success"`
	Error   string `json:"error"`
}

type LogoutResponse struct {
	Success bool   `json:"success" example:"true"`
	Error   string `json:"error,omitempty" example:"error"`
	Data    string `json:"data" example:"Logout success"`
}
