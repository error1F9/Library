package modules

import (
	astorage "Library/internal/modules/author/repository"
	bstorage "Library/internal/modules/book/repository"
	ustorage "Library/internal/modules/user/repository"
	"gorm.io/gorm"
)

type Repositories struct {
	User   ustorage.UserRepositoryInterface
	Author astorage.AuthorRepositoryInterface
	Book   bstorage.BookRepositoryInterface
}

func NewRepositories(db *gorm.DB) *Repositories {
	return &Repositories{
		Author: astorage.NewAuthorRepository(db),
		Book:   bstorage.NewBookRepository(db),
		User:   ustorage.NewUserRepository(db),
	}
}
