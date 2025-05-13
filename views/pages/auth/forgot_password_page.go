package auth

import (
	html "github.com/alexisbcz/libhtml"
	"github.com/alexisbcz/vauban/views/layouts"
)

func ForgotPasswordPage() html.Node {
	return layouts.AuthLayout(layouts.AuthLayoutProps{})()
}
