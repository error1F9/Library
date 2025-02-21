package service

import (
	"Library/internal/models"
	arepository "Library/internal/modules/author/repository"
	brepository "Library/internal/modules/book/repository"
	urepository "Library/internal/modules/user/repository"
	"Library/internal/pkg/cryptography"
	"Library/internal/token"
	"context"
	"database/sql"
	"errors"
	"github.com/go-chi/jwtauth/v5"
	"go.uber.org/zap"
)

type UserService struct {
	userRepository   urepository.UserRepositoryInterface
	bookRepository   brepository.BookRepositoryInterface
	authorRepository arepository.AuthorRepositoryInterface
	logger           *zap.Logger
	tokenService     *token.JWTTokenService
}

func NewUserService(userRepository urepository.UserRepositoryInterface, bookRepository brepository.BookRepositoryInterface, authorRepository arepository.AuthorRepositoryInterface, logger *zap.Logger, tokenService *token.JWTTokenService) *UserService {
	return &UserService{userRepository: userRepository, bookRepository: bookRepository, authorRepository: authorRepository, logger: logger, tokenService: tokenService}
}

func (s *UserService) CreateUser(in CreateUserIn) CreateUserOut {
	hashedPassword, err := cryptography.HashPassword(in.Password)
	if err != nil {
		s.logger.Error("Error hashing password", zap.Error(err))
		return CreateUserOut{0, err}
	}

	user := &models.User{
		Firstname: in.Firstname,
		Lastname:  in.Lastname,
		Email:     in.Email,
		Password:  hashedPassword,
	}
	id, err := s.userRepository.CreateUser(user)
	if err != nil {
		s.logger.Error("Error creating user", zap.Error(err))
	}
	return CreateUserOut{id, err}
}

func (s *UserService) RentBookByUser(in RentBookByUserIn) error {
	book, err := s.bookRepository.GetBookByID(in.BookID)
	if err != nil {
		s.logger.Error("Error getting book", zap.Error(err))
		return err
	}

	if book.RentedByID.Valid {
		s.logger.Error("Book has already rented", zap.Error(err))
		return errors.New("book has rented already")
	}

	user, err := s.userRepository.FindUserByID(in.UserID)
	if err != nil {
		s.logger.Error("Error finding user", zap.Error(err))
		return err
	}

	book.RentedByID = sql.Null[uint]{
		Valid: true,
		V:     user.ID,
	}
	author := book.Author
	author.Rating++
	if err = s.authorRepository.UpdateAuthor(&author); err != nil {
		s.logger.Error("Error updating author", zap.Error(err))
		return err
	}
	err = s.bookRepository.UpdateBook(book)
	if err != nil {
		s.logger.Error("Error updating book", zap.Error(err))
	}
	return err
}

func (s *UserService) ReturnBook(in ReturnBookIn) error {
	book, err := s.bookRepository.GetBookByID(in.BookID)
	if err != nil {
		s.logger.Error("Error getting book", zap.Error(err))
		return err
	}

	if !book.RentedByID.Valid {
		s.logger.Error("Book has not rented", zap.Error(err))
		return errors.New("book has not rented")
	}

	user, err := s.userRepository.FindUserByID(in.UserID)
	if err != nil {
		s.logger.Error("Error finding user", zap.Error(err))
		return err
	}

	if book.RentedByID.V != user.ID {
		s.logger.Error("User has not rented this book", zap.Error(err))
		return errors.New("user has not rented this book")
	}

	book.RentedByID = sql.Null[uint]{Valid: false}

	author := book.Author
	author.Rating--
	if err = s.authorRepository.UpdateAuthor(&author); err != nil {
		s.logger.Error("Error updating author", zap.Error(err))
		return err
	}

	err = s.bookRepository.ReturnBook(book)
	if err != nil {
		s.logger.Error("Error Returning book", zap.Error(err))
	}
	return err
}

func (s *UserService) GetAllUsers() GetAllUsersOut {
	users, err := s.userRepository.GetAllUsers()
	if err != nil {
		s.logger.Error("Error getting all users", zap.Error(err))
	}
	return GetAllUsersOut{users, err}
}

func (s *UserService) Login(in LoginIn) LoginOut {
	user, err := s.userRepository.GetUserByEmail(in.Email)
	if err != nil {
		s.logger.Error("Error getting user by email", zap.Error(err))
		return LoginOut{"", err}
	}

	if isValid := cryptography.CheckPassword(user.Password, in.Password); !isValid {
		s.logger.Error("Login", zap.Error(errors.New("invalid password")))
		return LoginOut{"", errors.New("invalid password")}
	}

	newToken, err := s.tokenService.GenerateToken(user.Email, user.ID)
	if err != nil {
		s.logger.Error("Error generating token", zap.Error(err))
		return LoginOut{"", err}
	}

	err = s.userRepository.Login(in.Email)
	if err != nil {
		s.logger.Error("Error logging in", zap.Error(err))
	}
	return LoginOut{newToken, err}
}

func (s *UserService) Logout(ctx context.Context) error {
	_, claims, _ := jwtauth.FromContext(ctx)
	email, ok := claims["email"].(string)
	if !ok {
		return errors.New("username not found")
	}
	err := s.userRepository.Logout(email)
	return err
}

func (s *UserService) Empty() (bool, error) {
	return s.userRepository.Empty()
}
