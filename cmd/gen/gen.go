package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"text/template"
)

// HTMLElement represents an HTML element with its attributes
type HTMLElement struct {
	Name       string
	Void       bool // True for void elements that don't need closing tags
	Attributes []string
}

// Template for generating element structs and functions
const elementTemplate = `
type {{.Name}} struct {
	*Tag
}

func {{title .Name}}(children ...html.Node) *{{.Name}} {
	return &{{.Name}}{New("{{.Name}}", {{.Void}}, children)}
}

{{range .Attributes}}
func (e *{{$.Name}}) {{title .}}(value string) *{{$.Name}} {
	e.attributes["{{.}}"] = value
	return e
}
{{end}}
`

// title capitalizes the first letter of a string
func title(s string) string {
	if s == "" {
		return ""
	}
	return strings.ToUpper(s[:1]) + s[1:]
}

func main() {
	// Path to the JSON file containing HTML elements data
	// You would need to download this file from the MDN repo or create it
	jsonFile := "html_elements.json"

	// Read JSON file
	data, err := ioutil.ReadFile(jsonFile)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		// If file doesn't exist, create example data
		elements := getHardcodedElements()
		generateCode(elements)
		return
	}

	// Parse JSON
	var elements []HTMLElement
	err = json.Unmarshal(data, &elements)
	if err != nil {
		fmt.Printf("Error parsing JSON: %v\n", err)
		return
	}

	generateCode(elements)
}

func generateCode(elements []HTMLElement) {
	// Create template with custom function
	tmpl := template.New("element")
	tmpl = tmpl.Funcs(template.FuncMap{"title": title})
	tmpl, err := tmpl.Parse(elementTemplate)
	if err != nil {
		fmt.Printf("Error parsing template: %v\n", err)
		return
	}

	// Create output file
	outputFile, err := os.Create("generated_html_elements.go")
	if err != nil {
		fmt.Printf("Error creating output file: %v\n", err)
		return
	}
	defer outputFile.Close()

	// Write package header
	fmt.Fprintln(outputFile, "package html")
	fmt.Fprintln(outputFile, "")
	fmt.Fprintln(outputFile, "import (")
	fmt.Fprintln(outputFile, `	"github.com/alexisbcz/vauban/pkg/html"`)
	fmt.Fprintln(outputFile, ")")

	// Generate code for each element
	for _, element := range elements {
		err = tmpl.Execute(outputFile, element)
		if err != nil {
			fmt.Printf("Error executing template for %s: %v\n", element.Name, err)
		}
	}

	fmt.Println("Code generation complete! Check generated_html_elements.go")
}

// getHardcodedElements returns a small set of hardcoded elements
// This is just an example - in practice you'd want to use a complete dataset
func getHardcodedElements() []HTMLElement {
	return []HTMLElement{
		{
			Name:       "div",
			Void:       false,
			Attributes: []string{"class", "id", "style"},
		},
		{
			Name:       "img",
			Void:       true,
			Attributes: []string{"src", "alt", "width", "height", "class", "id"},
		},
		// Add more elements as needed
	}
}
