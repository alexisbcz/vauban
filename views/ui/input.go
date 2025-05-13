package ui

import html "github.com/alexisbcz/libhtml"

type InputProps struct {
	Class       string
	Id          string
	Type        string
	Placeholder string
}

func Input(props InputProps) func(children ...html.Node) html.Node {
	return func(children ...html.Node) html.Node {
		return html.Input().
			Id(props.Id).Type(props.Type).Placeholder(props.Placeholder).
			Class("border-neutral-300 flex h-9 w-full rounded-md border bg-transparent px-3 py-1 text-base shadow-sm transition-colors file:border-0 file:bg-transparent file:text-sm file:font-medium file:text-foreground placeholder:text-muted-foreground focus-visible:outline-none focus-visible:ring-1 focus-visible:ring-neutral-500 disabled:cursor-not-allowed disabled:opacity-50 md:text-sm", props.Class).
			Children(children...)
	}
}
