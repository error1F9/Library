package models

import (
	"database/sql"
)

type Author struct {
	ID        uint   `json:"id" gorm:"primaryKey"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Books     []Book `json:"books,omitempty" gorm:"foreignKey:AuthorID"`
	Rating    int    `json:"rating"`
}

type Book struct {
	ID         uint           `json:"id" gorm:"primaryKey"`
	Title      string         `json:"title"`
	AuthorID   uint           `json:"author_id"`
	Author     Author         `json:"author" gorm:"foreignKey:AuthorID"`
	RentedByID sql.Null[uint] `json:"rented_by_id"`
	User       User           `json:"user" gorm:"foreignKey:RentedByID"`
}

type User struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	Firstname   string `json:"firstname gorm:not null"`
	Lastname    string `json:"lastname"`
	Email       string `json:"email,gorm:unique"`
	Password    string `json:"password,gorm:not null"`
	RentedBooks []Book `json:"rented_books" gorm:"foreignKey:RentedByID"`
	UserStatus  int    `json:"user_status"`
}
