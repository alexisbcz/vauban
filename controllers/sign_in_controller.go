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

	"github.com/alexisbcz/vauban/database/repositories"
	"github.com/alexisbcz/vauban/views/pages"
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
	return pages.SignInPage().Render(w)
}

func (c *SignInController) Handle(w http.ResponseWriter, r *http.Request) error {
	return nil
}
