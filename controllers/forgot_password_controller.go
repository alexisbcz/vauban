package controllers

import (
	"net/http"

	"github.com/alexisbcz/vauban/database/repositories"
	"github.com/alexisbcz/vauban/views/pages/auth"
)

type ForgotPasswordController struct {
	usersRepository repositories.UsersRepository
}

func NewForgotPasswordController(usersRepository repositories.UsersRepository) *ForgotPasswordController {
	return &ForgotPasswordController{
		usersRepository: usersRepository,
	}
}

func (c *ForgotPasswordController) Show(w http.ResponseWriter, r *http.Request) error {
	return auth.ForgotPasswordPage().Render(w)
}

func (c *ForgotPasswordController) Handle(w http.ResponseWriter, r *http.Request) error {
	return nil
}
