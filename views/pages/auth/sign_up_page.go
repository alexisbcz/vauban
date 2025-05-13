package auth

import (
	html "github.com/alexisbcz/libhtml"
	"github.com/alexisbcz/vauban/views/layouts"
	"github.com/alexisbcz/vauban/views/ui"
)

func SignUpPage() html.Node {
	return layouts.AuthLayout(layouts.AuthLayoutProps{
		Title:       "Get Started",
		Description: "Continue with your email address, and your password.",
	})(
		html.Form(
			ui.Field(
				ui.Label(ui.LabelProps{Text: "Email address", For: "email"})(),
				ui.Input(ui.InputProps{Type: "email", Id: "email", Placeholder: "john.doe@example.com"})(),
			),
			ui.Field(
				ui.Label(ui.LabelProps{Text: "Password", For: "password"})(),
				ui.Input(ui.InputProps{Type: "password", Id: "password", Placeholder: "················"})(),
			),
			ui.Button(ui.ButtonProps{Text: "Sign Up"})(),
			html.P(
				html.Text("Already have an account?"),
				html.A(html.Text("Sign in")).Href("/sign-in").
					Class("ml-1 cursor-pointer underline text-blue-700 hover:text-blue-600 transition-colors"),
			).Class("text-sm text-neutral-800 text-center"),
		).Class("flex flex-col gap-y-5"),
	)
}
