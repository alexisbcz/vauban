package controllers

import (
	"net/http"

	"github.com/alexisbcz/vauban/database/repositories"
	"github.com/alexisbcz/vauban/views/pages/auth"
)

type ResetPasswordController struct {
	usersRepository repositories.UsersRepository
}

func NewResetPasswordController(usersRepository repositories.UsersRepository) *ResetPasswordController {
	return &ResetPasswordController{
		usersRepository: usersRepository,
	}
}

func (c *ResetPasswordController) Show(w http.ResponseWriter, r *http.Request) error {
	return auth.ResetPasswordPage().Render(w)
}

func (c *ResetPasswordController) Handle(w http.ResponseWriter, r *http.Request) error {
	return nil
}
