package models

type UserDTO struct {
	ID          uint       `json:"id"`
	Firstname   string     `json:"firstname"`
	Lastname    string     `json:"lastname"`
	Email       string     `json:"email"`
	Password    string     `json:"password"`
	RentedBooks []*BookDTO `json:"rented_books,omitempty"`
	UserStatus  int        `json:"user_status,omitempty"`
}
