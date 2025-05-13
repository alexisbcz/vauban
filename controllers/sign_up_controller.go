package controllers

import (
	"net/http"

	"github.com/alexisbcz/vauban/database/models"
	"github.com/alexisbcz/vauban/database/repositories"
	"github.com/alexisbcz/vauban/views/pages/auth"
)

type SignUpController struct {
	usersRepository repositories.UsersRepository
}

func NewSignUpController(usersRepository repositories.UsersRepository) *SignUpController {
	return &SignUpController{
		usersRepository: usersRepository,
	}
}

func (c *SignUpController) Show(w http.ResponseWriter, r *http.Request) error {
	return auth.SignUpPage().Render(w)
}

func (c *SignUpController) Handle(w http.ResponseWriter, r *http.Request) error {
	if err := r.ParseForm(); err != nil {
		return err
	}

	if err := c.usersRepository.Store(r.Context(), &models.User{
		Email:    r.FormValue("email"),
		Password: r.FormValue("password"),
	}); err != nil {
		return err
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)

	return nil
}
