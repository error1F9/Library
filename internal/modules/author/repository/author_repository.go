package repository

import (
	"Library/internal/models"
	"gorm.io/gorm"
)

type AuthorRepository struct {
	db *gorm.DB
}

type AuthorRepositoryInterface interface {
	AddAuthor(author *models.Author) (*models.Author, error)
	GetAuthorByID(id uint) (*models.Author, error)
	UpdateAuthor(author *models.Author) error
	AddBookToAuthor(book *models.Book) error
	GetAllAuthors() ([]*models.Author, error)
	GetTopAuthors(limit int) ([]*models.Author, error)
	Empty() (bool, error)
}

func NewAuthorRepository(db *gorm.DB) *AuthorRepository {
	return &AuthorRepository{db: db}
}

func (r *AuthorRepository) AddAuthor(author *models.Author) (*models.Author, error) {
	err := r.db.Model(&models.Author{}).Create(&author).Error
	return author, err
}

func (r *AuthorRepository) GetAuthorByID(id uint) (*models.Author, error) {
	author := &models.Author{}
	err := r.db.Model(&models.Author{}).Preload("Books").First(author, id).Error
	return author, err
}

func (r *AuthorRepository) UpdateAuthor(author *models.Author) error {
	err := r.db.Model(&author).Save(&author).Error
	return err
}

func (r *AuthorRepository) AddBookToAuthor(book *models.Book) error {
	author, err := r.GetAuthorByID(book.AuthorID)
	if err != nil {
		return err
	}
	err = r.db.Model(&author).Association("Books").Append(&book)
	return err
}

func (r *AuthorRepository) GetAllAuthors() ([]*models.Author, error) {
	var authors []*models.Author
	err := r.db.Preload("Books").Find(&authors).Error
	return authors, err
}

func (r *AuthorRepository) GetTopAuthors(limit int) ([]*models.Author, error) {
	var authors []*models.Author
	err := r.db.Order("rating DESC").Limit(limit).Find(&authors).Error
	return authors, err
}

func (r *AuthorRepository) Empty() (bool, error) {
	var count int64
	err := r.db.Model(&models.Author{}).Limit(1).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count == 0, nil
}
