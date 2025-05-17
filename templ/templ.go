// Package templ provides the logic for parsing templates ang generate source code
package templ

import (
	"io"
	"text/template"
)

// Template definition
type Template struct {
	tmpl *template.Template
}

// Create a new template instance
func NewTemplate(templateSource io.Reader, name string) (*Template, error) {
	tmp := template.New(name)
	templData, err := io.ReadAll(templateSource)
	if err != nil {
		return nil, err
	}
	templ, err := tmp.Parse(string(templData))
	if err != nil {
		return nil, err
	}
	return &Template{
		tmpl: templ,
	}, nil
}

// Render the template into the outPut
func (t *Template) Render(outPut io.Writer, data any) error {
	return t.tmpl.Execute(outPut, data)
}
