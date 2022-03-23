package main

import (
	"html/template"
	"path/filepath"

	"github.com/esirk/snippet_box/pkg/models"
)

type templateData struct {
	Snippet *models.Snippet
	Snippets []*models.Snippet
}

func newTemplateCache(dir string)(map[string]*template.Template, error) {
	// map acting as cache
	cache := map[string]*template.Template{}

	// list of all templates to parse
	//home.page.tmpl
	files, err := filepath.Glob(filepath.Join(dir, "*.page.tmpl"))
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		name := filepath.Base(file)
		// parse the template file
		ts, err := template.ParseFiles(file)
		if err != nil {
			return nil, err
		}

		// Use ParseGlob method to add any 'layout' templates to the template set
		ts, err = ts.ParseGlob(filepath.Join(dir, "*.layout.tmpl"))
		if err != nil {
			return nil, err
		}

		// Use ParseGlob method to add any 'partial' templates to the template set
		ts, err = ts.ParseGlob(filepath.Join(dir, "*.partial.tmpl"))
		if err != nil {
			return nil, err
		}

		cache[name] = ts
	}
	return cache, nil
}
