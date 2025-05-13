package controllers

import (
	"net/http"

	"github.com/alexisbcz/vauban/database/repositories"
	"github.com/alexisbcz/vauban/views/pages/auth"
)

type SignInController struct {
	usersRepository repositories.UsersRepository
}

func NewSignInController(usersRepository repositories.UsersRepository) *SignInController {
	return &SignInController{
		usersRepository: usersRepository,
	}
}

func (c *SignInController) Show(w http.ResponseWriter, r *http.Request) error {
	return auth.SignInPage().Render(w)
}

func (c *SignInController) Handle(w http.ResponseWriter, r *http.Request) error {
	return nil
}
