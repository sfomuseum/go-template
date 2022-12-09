// Package text provides methods for loading text (.txt) templates with default functions
package text

import (
	"context"
	"fmt"
	"github.com/sfomuseum/go-template/funcs"
	"io/fs"
	"text/template"
)

// LoadTemplates loads text (.txt) from 't_fs' with default functions assigned.
func LoadTemplates(ctx context.Context, t_fs ...fs.FS) (*template.Template, error) {

	funcs := TemplatesFuncMap()
	t := template.New("text").Funcs(funcs)

	var err error

	for idx, f := range t_fs {

		t, err = t.ParseFS(f, "*.txt")

		if err != nil {
			return nil, fmt.Errorf("Failed to load templates from FS at offset %d, %w", idx, err)
		}
	}

	return t, nil
}

// TemplatesFuncMap() returns a `template.FuncMap` instance with default functions assigned.
func TemplatesFuncMap() template.FuncMap {

	return template.FuncMap{
		// For example: {{ if (IsAvailable "Account" .) }}
		"IsAvailable": funcs.IsAvailable,
	}
}
