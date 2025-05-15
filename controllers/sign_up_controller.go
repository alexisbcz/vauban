/*
Copyright 2025 Alexis Bouchez <alexbcz@proton.me>

This file is part of Vauban.

Vauban is free software: you can redistribute it and/or modify
it under the terms of the GNU Affero General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

Vauban is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License
along with Vauban. If not, see <https://www.gnu.org/licenses/>.

For commercial licensing options, please contact Alexis Bouchez at alexbcz@proton.me
*/
package controllers

import (
	"net/http"
	"time"

	"fmt"

	"github.com/alexisbcz/vauban/database/models"
	"github.com/alexisbcz/vauban/database/repositories"
	"github.com/alexisbcz/vauban/httperror"
	"github.com/alexisbcz/vauban/session"
	"github.com/alexisbcz/vauban/validator"
	"github.com/alexisbcz/vauban/validator/field"
	"github.com/alexisbcz/vauban/views/forms"
	"github.com/alexisbcz/vauban/views/pages"
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
	return pages.SignUpPage().Render(w)
}

func (c *SignUpController) Handle(w http.ResponseWriter, r *http.Request) error {
	// Parse the incoming request form body
	if err := r.ParseForm(); err != nil {
		return httperror.BadRequest("cannot parse form")
	}

	// Define the form validator
	v := validator.New(
		field.New("firstName").MinLength(1).MaxLength(50),
		field.New("lastName").MinLength(1).MaxLength(50),
		field.
			New("email").Email().MaxLength(255).
			Unique(func(value string) (ok bool) {
				_, err := c.usersRepository.GetByEmail(r.Context(), value)
				return err == repositories.ErrUserNotFound
			}),
		field.New("password").MinLength(8).MaxLength(255),
	)

	// Validate the incoming request form body
	valid, errors := v.ValidateForm(r.Form)
	if !valid {
		return forms.SignUpForm(errors).Render(w)
	}

	// Define the new user record
	user := &models.User{
		FirstName: r.FormValue("firstName"),
		LastName:  r.FormValue("lastName"),
		Email:     r.FormValue("email"),
		Password:  r.FormValue("password"),
	}

	// Store a new user record in the database
	if err := c.usersRepository.Store(r.Context(), user); err != nil {
		return httperror.InternalServerError("failed to store the user record in the database")
	}

	// Authenticate the user
	session.SetSignedCookie(w, session.AUTH_COOKIE_NAME, fmt.Sprint(user.ID), int(time.Hour*24*7))

	// Redirect using HTMX
	w.Header().Set("HX-Redirect", "/")

	return nil
}
