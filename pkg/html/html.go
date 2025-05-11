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

package html

import (
	"fmt"
	"html"
	"io"
)

// Attribute represents HTML element attributes
type Attribute map[string]string

// Node interface defines components that can render themselves
type Node interface {
	Render(w io.Writer) error
}

// document represents an HTML document with its structure
type document struct {
	children []Node
}

// Render implements Node for document, rendering a complete HTML document
func (d *document) Render(w io.Writer) (err error) {
	if _, err := w.Write([]byte("<!DOCTYPE html>")); err != nil {
		return err
	}
	for _, child := range d.children {
		if err := child.Render(w); err != nil {
			return err
		}
	}
	return nil
}

// Document creates a complete HTML document with the given children
func Document(children ...Node) Node {
	return &document{
		children: children,
	}
}

// text renders escaped text content
type text struct {
	content string
}

// Text creates a node that renders HTML-escaped text content
func Text(content string) Node {
	return &text{content: content}
}

// Textf creates a node that renders formatted HTML-escaped text
func Textf(format string, args ...any) Node {
	return &text{content: fmt.Sprintf(format, args...)}
}

// Render implements Node.Render for text
func (t *text) Render(w io.Writer) error {
	_, err := io.WriteString(w, html.EscapeString(t.content))
	return err
}

// raw renders content as-is without escaping
type raw struct {
	content string
}

// Raw creates a node that renders raw HTML content
func Raw(content string) Node {
	return &raw{content: content}
}

// Rawf creates a node that renders formatted raw HTML content
func Rawf(format string, args ...any) Node {
	return &raw{content: fmt.Sprintf(format, args...)}
}

// Render implements Node.Render for raw
func (r *raw) Render(w io.Writer) error {
	_, err := io.WriteString(w, r.content)
	return err
}

// if_ conditionally renders content based on a condition
type if_ struct {
	condition bool
	then      Node
}

// If conditionally renders content when condition is true
func If(condition bool, then Node) Node {
	return &if_{
		condition: condition,
		then:      then,
	}
}

// ifFunc is a lazy conditional renderer that only evaluates its content when true
type ifFunc struct {
	condition bool
	thenFn    func() Node
}

// IfFunc conditionally renders content when condition is true
// Uses a callback function to avoid evaluating the content when condition is false
func IfFunc(condition bool, thenFn func() Node) Node {
	return &ifFunc{
		condition: condition,
		thenFn:    thenFn,
	}
}

// Render implements Node.Render for ifFunc
func (i *ifFunc) Render(w io.Writer) error {
	if i.condition && i.thenFn != nil {
		node := i.thenFn()
		if node != nil {
			return node.Render(w)
		}
	}
	return nil
}

// Render implements Node.Render for if_
func (i *if_) Render(w io.Writer) error {
	if i.condition {
		return i.then.Render(w)
	}
	return nil
}

// ifElse conditionally renders one of two contents based on a condition
type ifElse struct {
	condition bool
	then      Node
	else_     Node
}

// IfElse conditionally renders one of two nodes based on condition
func IfElse(condition bool, then Node, else_ Node) Node {
	return &ifElse{
		condition: condition,
		then:      then,
		else_:     else_,
	}
}

// Render implements Node.Render for ifElse
func (ie *ifElse) Render(w io.Writer) error {
	if ie.condition {
		return ie.then.Render(w)
	}
	return ie.else_.Render(w)
}

// ifElseFunc is a lazy conditional renderer that only evaluates its content when true
type ifElseFunc struct {
	condition bool
	thenFn    func() Node
	elseFn    func() Node
}

// IfElseFunc conditionally renders content when condition is true
// Uses a callback function to avoid evaluating the content when condition is false
func IfElseFunc(condition bool, thenFn, elseFn func() Node) Node {
	return &ifElseFunc{
		condition: condition,
		thenFn:    thenFn,
		elseFn:    elseFn,
	}
}

// Render implements Node.Render for ifElseFunc
func (i *ifElseFunc) Render(w io.Writer) error {
	if i.condition && i.thenFn != nil {
		node := i.thenFn()
		if node != nil {
			return node.Render(w)
		}
		return nil
	}
	return i.elseFn().Render(w)
}

// map_ renders a collection of items using a mapping function
type map_[T any] struct {
	items     []T
	transform func(item T) Node
}

// Map renders a collection of items using a transform function
func Map[T any](items []T, transform func(item T) Node) Node {
	return &map_[T]{
		items:     items,
		transform: transform,
	}
}

// Render implements Node.Render for map_
func (m *map_[T]) Render(w io.Writer) error {
	for _, item := range m.items {
		node := m.transform(item)
		if err := node.Render(w); err != nil {
			return err
		}
	}
	return nil
}

// group represents a collection of nodes with no root element
type group struct {
	children []Node
}

// Group combines multiple nodes without a wrapper element
func Group(children ...Node) Node {
	return &group{children: children}
}

// Render implements Node.Render for group
func (g *group) Render(w io.Writer) error {
	for _, child := range g.children {
		if err := child.Render(w); err != nil {
			return err
		}
	}
	return nil
}
