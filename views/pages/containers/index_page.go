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
package containers

import (
	"strings"

	html "github.com/alexisbcz/libhtml"
	"github.com/alexisbcz/vauban/views/layouts"
	"github.com/alexisbcz/vauban/views/ui"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/swarm"
)

func IndexPage(containers []container.Summary, services []swarm.Service) html.Node {
	return layouts.DashboardLayout(layouts.DashboardLayoutProps{
		Title: "Containers",
	})(
		html.Main(
			html.Map(containers, containerCard),
		).Class("mx-auto max-w-5xl my-12 grid grid-cols-4 gap-4"),
	)
}

func containerCard(container container.Summary) html.Node {
	return ui.Card(ui.CardProps{})(
		html.Span(html.Text("📦")),
		html.P(html.Text(strings.Join(container.Names, ""))),
	)
}
