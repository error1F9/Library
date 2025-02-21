package service

import (
	"Library/internal/models"
	"Library/internal/modules/author/repository"
	"go.uber.org/zap"
)

type AuthorService struct {
	authorRepository repository.AuthorRepositoryInterface
	logger           *zap.Logger
}

func NewAuthorService(authorRepository repository.AuthorRepositoryInterface, logger *zap.Logger) *AuthorService {
	return &AuthorService{authorRepository: authorRepository, logger: logger}
}

func (s *AuthorService) AddAuthor(in AddAuthorIn) AddAuthorOut {
	author, err := s.authorRepository.AddAuthor(&models.Author{Firstname: in.Firstname, Lastname: in.Lastname})
	if err != nil {
		s.logger.Error("Add Author: ", zap.Error(err))
	}
	return AddAuthorOut{
		Error:  err,
		Author: author,
	}
}

func (s *AuthorService) GetAllAuthors() GetAllAuthorsOut {
	authors, err := s.authorRepository.GetAllAuthors()
	if err != nil {
		s.logger.Error("Get All Authors: ", zap.Error(err))
	}
	return GetAllAuthorsOut{
		Authors: authors,
		Error:   err,
	}
}

func (s *AuthorService) GetTopAuthors(limit int) GetAllAuthorsOut {
	authors, err := s.authorRepository.GetTopAuthors(limit)
	if err != nil {
		s.logger.Error("Get Top Authors: ", zap.Error(err))
	}
	return GetAllAuthorsOut{
		Authors: authors,
		Error:   err,
	}
}

func (s *AuthorService) Empty() (bool, error) {
	return s.authorRepository.Empty()
}
