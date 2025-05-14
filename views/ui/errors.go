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

import (
	"fmt"
	"unicode"

	html "github.com/alexisbcz/libhtml"
	"github.com/alexisbcz/vauban/validator"
)

func Errors(key string, errorsData validator.Errors) html.Node {
	fmt.Println(errorsData)
	errors := errorsData[key]

	if len(errorsData[key]) == 0 {
		return nil
	}

	return html.Div(
		html.Map(errors, func(err string) html.Node {
			return html.
				P(html.Textf("→ %s %s", capitalize(key), err)).
				Class("text-sm text-red-600")
		}),
	)
}

func capitalize(s string) string {
	r := []rune(s)
	r[0] = unicode.ToUpper(r[0])
	return string(r)
}
