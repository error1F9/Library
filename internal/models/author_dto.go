package models

type AuthorDTO struct {
	ID        uint       `json:"id"`
	Firstname string     `json:"firstname"`
	Lastname  string     `json:"lastname"`
	Books     []*BookDTO `json:"books,omitempty"`
	Rating    int        `json:"rating"`
}
