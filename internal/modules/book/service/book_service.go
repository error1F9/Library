package service

import (
	"Library/internal/models"
	arepository "Library/internal/modules/author/repository"
	brepository "Library/internal/modules/book/repository"
	"go.uber.org/zap"
)

type BookService struct {
	bookRepository   brepository.BookRepositoryInterface
	authorRepository arepository.AuthorRepositoryInterface
	logger           *zap.Logger
}

func NewBookService(bookRepository brepository.BookRepositoryInterface, authorRepository arepository.AuthorRepositoryInterface, logger *zap.Logger) *BookService {
	return &BookService{bookRepository: bookRepository, authorRepository: authorRepository, logger: logger}
}

func (s *BookService) AddBook(in AddBookIn) AddBookOut {
	book, err := s.bookRepository.AddBook(&models.Book{Title: in.BookTitle, AuthorID: in.AuthorID})
	if err != nil {
		s.logger.Error("Add book: ", zap.Error(err))
		return AddBookOut{Error: err}
	}

	author, err := s.authorRepository.GetAuthorByID(in.AuthorID)
	if err != nil {
		s.logger.Error("Get Author By ID: ", zap.Error(err))
		return AddBookOut{Error: err}
	}

	book.Author = *author

	err = s.bookRepository.UpdateBook(book)
	if err != nil {
		s.logger.Error("Update Book: ", zap.Error(err))
	}
	return AddBookOut{
		Error:  err,
		Book:   book,
		Author: author,
	}
}

func (s *BookService) GetAllBooks() GetAllBooksOut {
	books, err := s.bookRepository.GetAllBooks()
	if err != nil {
		s.logger.Error("GetAllBooks: ", zap.Error(err))
	}
	return GetAllBooksOut{
		Books: books,
		Error: err,
	}
}

func (s *BookService) Empty() (bool, error) {
	return s.bookRepository.Empty()
}
