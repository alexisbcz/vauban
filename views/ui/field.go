package ui

import html "github.com/alexisbcz/libhtml"

func Field(children ...html.Node) html.Node {
	return html.Div().
		Class("flex flex-col gap-y-2").
		Children(children...)
}
