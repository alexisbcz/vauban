package ui

import html "github.com/alexisbcz/libhtml"

type LabelProps struct {
	Class string
	Text  string
	For   string
}

func Label(props LabelProps) func(children ...html.Node) html.Node {
	return func(children ...html.Node) html.Node {
		return html.Label().
			For(props.For).
			Class("text-sm font-medium leading-none peer-disabled:cursor-not-allowed peer-disabled:opacity-70", props.Class).
			Children(html.Text(props.Text), html.Group(children...))
	}
}
