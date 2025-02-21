package modules

import (
	aservice "Library/internal/modules/author/service"
	bservice "Library/internal/modules/book/service"
	uservice "Library/internal/modules/user/service"
	"Library/internal/pkg/generate"
	"Library/internal/token"
	"fmt"
	"go.uber.org/zap"
)

type Services struct {
	User   uservice.UserServiceInterface
	Book   bservice.BookServiceInterface
	Author aservice.AuthorServiceInterface
}

func NewServices(repositories *Repositories, logger *zap.Logger, token *token.JWTTokenService) *Services {
	bookService := bservice.NewBookService(repositories.Book, repositories.Author, logger)
	authorService := aservice.NewAuthorService(repositories.Author, logger)
	userService := uservice.NewUserService(repositories.User, repositories.Book, repositories.Author, logger, token)
	return &Services{
		Book:   bookService,
		Author: authorService,
		User:   userService,
	}
}

func (s *Services) IsEmpty() (bool, error) {
	bookEmpty, err := s.Book.Empty()
	if err != nil {
		return false, fmt.Errorf("failed to check if books table is empty: %w", err)
	}

	authorEmpty, err := s.Author.Empty()
	if err != nil {
		return false, fmt.Errorf("failed to check if authors table is empty: %w", err)
	}

	userEmpty, err := s.User.Empty()
	if err != nil {
		return false, fmt.Errorf("failed to check if users table is empty: %w", err)
	}

	return bookEmpty && authorEmpty && userEmpty, nil
}

func (s *Services) FillBD(nAuthors int, nBooks int, nUsers int) error {
	firstname, lastname := generate.Authors(nAuthors)
	authorIDs := make([]uint, nAuthors)

	for i := 0; i < nAuthors; i++ {
		out := s.Author.AddAuthor(aservice.AddAuthorIn{
			Firstname: firstname[i],
			Lastname:  lastname[i],
		})
		if out.Error != nil {
			return out.Error
		}
		authorIDs[i] = out.Author.ID
	}

	bookTitles := generate.Books(nBooks)
	for i := 0; i < nBooks; i++ {
		authorID := authorIDs[i%nAuthors]
		out := s.Book.AddBook(bservice.AddBookIn{
			BookTitle: bookTitles[i],
			AuthorID:  authorID,
		})
		if out.Error != nil {
			return out.Error
		}
	}

	users := generate.Users(nUsers)
	for i := 0; i < nUsers; i++ {
		out := s.User.CreateUser(uservice.CreateUserIn{
			Firstname: users[i].Firstname,
			Lastname:  users[i].Lastname,
			Email:     users[i].Email,
			Password:  users[i].Password,
		})
		if out.Error != nil {
			return out.Error
		}
	}

	return nil
}
