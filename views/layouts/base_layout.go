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

import html "github.com/alexisbcz/libhtml"

type BaseLayoutProps struct {
	Title string
}

func BaseLayout(props BaseLayoutProps) func(children ...html.Node) html.Node {
	return func(children ...html.Node) html.Node {
		return html.Document(
			html.HTML(
				html.Head(
					html.Meta().Charset("utf-8"),
					html.Meta().Name("viewport").Content("width=device-width, initial-scale=1"),
					html.Link().Rel("stylesheet").Href("/styles.css"),
					html.Script().Type("module").Src("https://cdn.jsdelivr.net/gh/starfederation/datastar@v1.0.0-beta.11/bundles/datastar.js"),
					html.Title(html.Text(props.Title)),
				),
				html.Body(children...),
			).Lang("en").Attribute("data-on-load", "@get('/hotreload', {retryInterval: 100})"),
		)
	}
}
