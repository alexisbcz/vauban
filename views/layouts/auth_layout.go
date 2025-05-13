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
package layouts

import (
	html "github.com/alexisbcz/libhtml"
	"github.com/alexisbcz/vauban/views/ui"
)

type AuthLayoutProps struct {
	Title       string
	Description string
}

func AuthLayout(props AuthLayoutProps) func(children ...html.Node) html.Node {
	return func(children ...html.Node) html.Node {
		return BaseLayout(BaseLayoutProps{
			Title: props.Title,
			Class: "bg-neutral-100 flex min-h-svh flex-col items-center justify-center gap-6 p-6 md:p-10",
		})(
			html.Main(
				ui.Card(ui.CardProps{Class: "w-full !py-6"})(
					html.H1(html.Text(props.Title)).Class("text-xl font-semibold text-center"),
					html.H2(html.Text(props.Description)).Class("mt-2 text-center text-neutral-500 text-sm"),
					html.Div(children...).Class("mt-4"),
				),
			).Class("flex w-full max-w-sm flex-col gap-6"),
		)
	}
}
