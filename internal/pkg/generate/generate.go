package generate

import (
	"Library/internal/models"
	"github.com/brianvoe/gofakeit/v7"
)

func Books(n int) []string {
	BookTitles := make([]string, n)
	for i := 0; i < n; i++ {
		BookTitles[i] = gofakeit.BookTitle()
	}
	return BookTitles
}

func Authors(n int) ([]string, []string) {
	AuthorsFirstNames := make([]string, n)
	AuthorsLastNames := make([]string, n)
	for i := 0; i < n; i++ {
		AuthorsFirstNames[i] = gofakeit.FirstName()
		AuthorsLastNames[i] = gofakeit.FirstName()
	}
	return AuthorsFirstNames, AuthorsLastNames
}

func Users(n int) []*models.UserDTO {
	users := make([]*models.UserDTO, n)
	for i := 0; i < n; i++ {
		users[i] = &models.UserDTO{}
		users[i].Firstname = gofakeit.FirstName()
		users[i].Lastname = gofakeit.LastName()
		users[i].Email = gofakeit.Email()
		users[i].Password = gofakeit.Password(true, true, true, true, false, 12)
	}
	return users
}
