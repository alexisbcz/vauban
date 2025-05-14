/* Copyright 2025 Alexis Bouchez <alexbcz@proton.me>  This file is part of Vauban.  Vauban is free software: you can redistribute it and/or modify it under the terms of the GNU Affero General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.  Vauban is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU Affero General Public License for more details.  You should have received a copy of the GNU Affero General Public License along with Vauban. If not, see <https://www.gnu.org/licenses/>.  For commercial licensing options, please contact Alexis Bouchez at alexbcz@proton.me */
package layouts

import html "github.com/alexisbcz/libhtml"

type BaseLayoutProps struct {
	Title string
	Class string
}

func BaseLayout(props BaseLayoutProps) func(children ...html.Node) html.Node {
	return func(children ...html.Node) html.Node {
		return html.Document(
			html.HTML(
				html.Head(
					html.Meta().Charset("utf-8"),
					html.Meta().Name("viewport").Content("width=device-width, initial-scale=1"),
					html.Link().Rel("stylesheet").Href("/styles.css"),
					html.Link().Rel("preconnect").Href("https://fonts.googleapis.com"),
					html.Link().Rel("preconnect").Href("https://fonts.gstatic.com").Attribute("crossorigin", ""),
					html.Link().Rel("stylesheet").Href("https://fonts.googleapis.com/css2?family=Inter:wght@100..900&family=Instrument+Serif:ital@0;1&display=swap"),
					html.Script().Type("module").Src("https://unpkg.com/htmx.org@2.0.4"),
					html.Title(html.Text(props.Title)),

					// Add hot reload script
					html.Script(html.Raw(`
						document.addEventListener('DOMContentLoaded', function() {
							function setupEventSource() {
								console.log('Connecting to hotreload SSE...');
								const eventSource = new EventSource('/hotreload');
								
								eventSource.addEventListener('open', function() {
									console.log('SSE connection established');
								});
								
								eventSource.addEventListener('ping', function(e) {
									console.log('Ping received:', e.data);
								});
								
								eventSource.addEventListener('reload', function() {
									console.log('Reload signal received');
									window.location.reload();
								});
								
								eventSource.addEventListener('error', function(e) {
									console.error('SSE error:', e);
									eventSource.close();
									
									// Try to reconnect after a delay
									setTimeout(function() {
										console.log('Attempting to reconnect...');
										setupEventSource();
									}, 2000);
								});
							}
							
							// Initial setup
							setupEventSource();
						});
					`)),
				),
				html.Body(children...).ClassIf(props.Class != "", props.Class),
			).Lang("en"),
		)
	}
}
