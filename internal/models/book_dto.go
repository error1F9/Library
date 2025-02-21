package models

type BookDTO struct {
	ID       uint       `json:"id"`
	Title    string     `json:"title"`
	Author   *AuthorDTO `json:"author,omitempty"`
	RentedBy *UserDTO   `json:"rented,omitempty"`
}
