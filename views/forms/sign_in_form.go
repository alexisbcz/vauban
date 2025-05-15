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
package forms

import (
	html "github.com/alexisbcz/libhtml"
	"github.com/alexisbcz/vauban/validator"
	"github.com/alexisbcz/vauban/views/ui"
)

func SignInForm(errors validator.Errors) html.Node {
	return html.Form(
		ui.Field(
			ui.Label(ui.LabelProps{Text: "Email address", For: "email"})(),
			ui.Input(ui.InputProps{Type: "email", Id: "email", Placeholder: "john.doe@example.com"})(),
			ui.Errors("email", errors),
		),
		ui.Field(
			html.Div(
				ui.Label(ui.LabelProps{Text: "Password", For: "password"})(),
				html.A(html.Text("Forgot your password?")).Href("/forgot-password").
					Class("ml-1 cursor-pointer text-sm underline text-blue-700 hover:text-blue-600 transition-colors"),
			).Class("flex flex-wrap gap-2 items-center justify-between"),
			ui.Input(ui.InputProps{Type: "password", Id: "password", Placeholder: "················"})(),
		),
		ui.Button(ui.ButtonProps{Text: "Sign In", Type: "submit"}),
		html.P(
			html.Text("Not signed up (yet)?"),
			html.A(html.Text("Sign up")).Href("/sign-up").
				Class("ml-1 cursor-pointer underline text-blue-700 hover:text-blue-600 transition-colors"),
		).Class("text-sm text-neutral-800 text-center"),
	).
		Id("sign-in-form").
		Attribute("hx-post", "/sign-in/").
		Attribute("hx-swap", "outerHTML").
		Attribute("hx-target", "#sign-in-form").
		Class("flex flex-col gap-y-5")
}
