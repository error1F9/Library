package repository

import (
	"Library/internal/models"
	"errors"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

type UserRepositoryInterface interface {
	CreateUser(user *models.User) (uint, error)
	FindUserByID(id uint) (*models.User, error)
	GetAllUsers() ([]*models.User, error)
	GetUserByEmail(email string) (*models.User, error)
	Login(email string) error
	Logout(email string) error
	Empty() (bool, error)
	//UpdateUser(user *models.User) (*models.User, error)
	//DeleteUser(user *models.User) error
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) CreateUser(user *models.User) (uint, error) {
	err := r.db.Model(&models.User{}).Create(&user).Error
	return user.ID, err
}

func (r *UserRepository) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	err := r.db.Table("users").Where("email = ?", email).First(&user).Error
	return &user, err
}

func (r *UserRepository) FindUserByID(id uint) (*models.User, error) {
	var user models.User
	if err := r.db.First(&user, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user does not exist")
		}
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) GetAllUsers() ([]*models.User, error) {
	var users []*models.User
	err := r.db.Find(&users).Error
	return users, err
}

func (r *UserRepository) Login(email string) error {
	err := r.db.Table("users").Where("email = ?", email).Update("user_status", 1).Error
	return err
}

func (r *UserRepository) Logout(email string) error {
	err := r.db.Table("users").Where("email = ?", email).Update("user_status", 0).Error
	return err
}

func (r *UserRepository) Empty() (bool, error) {
	var count int64
	err := r.db.Model(&models.User{}).Limit(1).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count == 0, nil
}

//func (r *UserRepository) UpdateUser(user *models.User) (*models.User, error) {
//	err := r.db.Model(&models.User{}).Save(&user).Error
//	return user, err
//}
//
//func (r *UserRepository) DeleteUser(user *models.User) error {
//	err := r.db.Model(&models.User{}).Delete(&user).Error
//	return err
//}
