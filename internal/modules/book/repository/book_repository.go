package repository

import (
	"Library/internal/models"
	"errors"
	"gorm.io/gorm"
)

type BookRepository struct {
	db *gorm.DB
}

type BookRepositoryInterface interface {
	AddBook(book *models.Book) (*models.Book, error)
	GetAllBooks() ([]*models.Book, error)
	UpdateBook(book *models.Book) error
	GetBookByID(id uint) (*models.Book, error)
	ReturnBook(book *models.Book) error
	Empty() (bool, error)
}

func NewBookRepository(db *gorm.DB) *BookRepository {
	return &BookRepository{db: db}
}

func (r *BookRepository) AddBook(book *models.Book) (*models.Book, error) {
	err := r.db.Model(&models.Book{}).Create(&book).Error
	return book, err
}

func (r *BookRepository) GetBookByID(id uint) (*models.Book, error) {
	var book models.Book
	if err := r.db.Preload("Author").Preload("User").First(&book, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("book not found")
		}
		return nil, err
	}
	return &book, nil
}

func (r *BookRepository) UpdateBook(book *models.Book) error {
	err := r.db.Save(&book).Error
	return err
}

func (r *BookRepository) ReturnBook(book *models.Book) error {
	err := r.db.Model(&models.Book{}).Where("id = ?", book.ID).Update("rented_by_id", nil).Error
	return err
}

func (r *BookRepository) GetAllBooks() ([]*models.Book, error) {
	var books []*models.Book
	err := r.db.Preload("Author").Preload("User").Find(&books).Error
	return books, err
}

func (r *BookRepository) Empty() (bool, error) {
	var count int64
	err := r.db.Model(&models.Book{}).Limit(1).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count == 0, nil
}
