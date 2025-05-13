package auth

import (
	html "github.com/alexisbcz/libhtml"
	"github.com/alexisbcz/vauban/views/layouts"
	"github.com/alexisbcz/vauban/views/ui"
)

func SignInPage() html.Node {
	return layouts.AuthLayout(layouts.AuthLayoutProps{
		Title:       "Welcome back",
		Description: "Continue with your email address, and your password.",
	})(
		html.Form(
			ui.Field(
				ui.Label(ui.LabelProps{Text: "Email address", For: "email"})(),
				ui.Input(ui.InputProps{Type: "email", Id: "email", Placeholder: "john.doe@example.com"})(),
			),
			ui.Field(
				html.Div(
					ui.Label(ui.LabelProps{Text: "Password", For: "password"})(),
					html.A(html.Text("Forgot your password?")).Href("/forgot-password").
						Class("ml-1 cursor-pointer text-sm underline text-blue-700 hover:text-blue-600 transition-colors"),
				).Class("flex flex-wrap gap-2 items-center justify-between"),
				ui.Input(ui.InputProps{Type: "password", Id: "password", Placeholder: "················"})(),
			),
			ui.Button(ui.ButtonProps{Text: "Sign In"})(),
			html.P(
				html.Text("Not signed up (yet)?"),
				html.A(html.Text("Sign up")).Href("/sign-up").
					Class("ml-1 cursor-pointer underline text-blue-700 hover:text-blue-600 transition-colors"),
			).Class("text-sm text-neutral-800 text-center"),
		).Class("flex flex-col gap-y-5"),
	)
}
