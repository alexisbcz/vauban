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

package tag

import (
	"fmt"
	"io"

	"github.com/alexisbcz/vauban/pkg/html"
)

type Tag struct {
	name       string
	ending     bool
	children   []html.Node
	attributes html.Attribute
}

func New(name string, ending bool, children []html.Node) *Tag {
	return &Tag{
		name:       name,
		ending:     ending,
		attributes: make(html.Attribute),
		children:   children,
	}
}

func (e *Tag) Children(children ...html.Node) *Tag {
	e.children = children
	return e
}

// Render implements Node.
func (e *Tag) Render(w io.Writer) error {
	if _, err := fmt.Fprintf(w, "<%s", e.name); err != nil {
		return err
	}

	// Render attributes
	for key, value := range e.attributes {
		if _, err := fmt.Fprintf(w, " %s=\"%s\"", key, value); err != nil {
			return err
		}
	}

	if e.ending {
		_, err := w.Write([]byte("/>"))
		return err
	}

	// Write closing bracket for opening tag
	if _, err := w.Write([]byte(">")); err != nil {
		return err
	}

	// Render all children
	for _, child := range e.children {
		if err := child.Render(w); err != nil {
			return err
		}
	}

	// Write closing tag
	_, err := fmt.Fprintf(w, "</%s>", e.name)
	return err
}

func (t *Tag) Attribute(key, value string) *Tag {
	t.attributes[key] = value
	return t
}
