// Package text provides methods for loading text (.txt) templates with default functions
package text

import (
	"context"
	"github.com/sfomuseum/go-template/funcs"
	"text/template"
	"io/fs"
)

// LoadTemplates loads text (.txt) from 't_fs' with default functions assigned.
func LoadTemplates(ctx context.Context, t_fs fs.FS) (*template.Template, error) {

	funcs := TemplatesFuncMap()
	t := template.New("text").Funcs(funcs)
	return t.ParseFS(t_fs, "*.txt")
}

// TemplatesFuncMap() returns a `template.FuncMap` instance with default functions assigned.
func TemplatesFuncMap() template.FuncMap {

	return template.FuncMap{
		// For example: {{ if (IsAvailable "Account" .) }}
		"IsAvailable": funcs.IsAvailable,
	}
}
