package router

import (
	"Library/internal/modules"
	"Library/internal/token"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/jwtauth/v5"
	httpSwagger "github.com/swaggo/http-swagger/v2"
)

func NewApiRouter(c *modules.Controllers, t *token.JWTTokenService) *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Route("/library", func(r chi.Router) {
		r.Route("/user", func(r chi.Router) {
			r.Group(func(r chi.Router) {
				r.Use(jwtauth.Verifier(t.GetJWTAuth()))
				r.Use(jwtauth.Authenticator(t.GetJWTAuth()))
				r.Post("/rent", c.User.RentBookByUser)
				r.Post("/return", c.User.ReturnBook)
				r.Post("/logout", c.User.Logout)
			})
			r.Post("/login", c.User.Login)
			r.Get("/all", c.User.GetAllUsers)
			r.Post("/create", c.User.CreateUser)
		})
		r.Route("/authors", func(r chi.Router) {
			r.Get("/all", c.Author.GetAllAuthors)
			r.Get("/top", c.Author.GetTopAuthors)
			r.Post("/add", c.Author.AddAuthor)
		})
		r.Route("/books", func(r chi.Router) {
			r.Get("/all", c.Book.GetAllBooks)
			r.Post("/add", c.Book.AddBook)
		})

	})

	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("/swagger/doc.json"),
	))
	return r
}
