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
			Id(props.Id).
			Name(props.Id).
			Type(props.Type).
			Placeholder(props.Placeholder).
			Class("border-neutral-300 flex h-9 w-full rounded-md border bg-transparent px-3 py-1 text-base shadow-sm transition-colors file:border-0 file:bg-transparent file:text-sm file:font-medium file:text-foreground placeholder:text-muted-foreground focus-visible:outline-none focus-visible:ring-1 focus-visible:ring-neutral-500 disabled:cursor-not-allowed disabled:opacity-50 md:text-sm", props.Class).
			Children(children...)
	}
}
