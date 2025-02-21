package modules

import (
	acontroller "Library/internal/modules/author/controller"
	bcontroller "Library/internal/modules/book/controller"
	ucontroller "Library/internal/modules/user/controller"
	"Library/internal/responder"
	"github.com/ptflp/godecoder"
)

type Controllers struct {
	User   ucontroller.UserControllerInterface
	Book   bcontroller.BookControllerInterface
	Author acontroller.AuthorControllerInterface
}

func NewControllers(services *Services, decoder godecoder.Decoder, responder responder.Responder) *Controllers {
	bookController := bcontroller.NewBookController(services.Book, decoder, responder)
	authorController := acontroller.NewAuthorController(services.Author, decoder, responder)
	userController := ucontroller.NewUserController(services.User, decoder, responder)

	return &Controllers{
		Book:   bookController,
		Author: authorController,
		User:   userController,
	}
}
