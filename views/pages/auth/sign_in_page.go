package auth

import (
	html "github.com/alexisbcz/libhtml"
	"github.com/alexisbcz/vauban/views/layouts"
)

func SignInPage() html.Node {
	return layouts.AuthLayout(layouts.AuthLayoutProps{})()
}
