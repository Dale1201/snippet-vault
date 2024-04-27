package main

import (
	"html/template"
	"path/filepath"

	"github.com/Dale1201/snippet-vault/pkg/models"
)

type templateData struct {
	CurrentYear int
	Snippet *models.Snippet
	Snippets []*models.Snippet
}

func newTemplateCache(dir string) (map[string]*template.Template, error) {
	// Initialize a new map to act as the cache. 
	cache := map[string]*template.Template{}

	// Use the filepath.Glob function to get a slice of all filepaths with 
    // the extension '.page.tmpl'. This essentially gives us a slice of all the 
    // 'page' templates for the application. 
	pages, err := filepath.Glob(filepath.Join(dir, "*.page.tmpl"))
	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		// Extract the file name (like 'home.page.tmpl') from the full file pat
        // and assign it to the name variable. 
		name := filepath.Base(page)

		// Parse the page template file in to a template set. 
		ts, err := template.ParseFiles(page)
		if err != nil {
			return nil, err
		}

        // Use the ParseGlob method to add any 'layout' templates to the 
        // template set (in our case, it's just the 'base' layout at the 
        // moment). 
		ts, err = ts.ParseGlob(filepath.Join(dir, "*.layout.tmpl"))
		if err != nil {
			return nil, err
		}

		// Use the ParseGlob method to add any 'partial' templates to the 
        // template set (in our case, it's just the 'footer' partial at the 
        // moment). 
		ts, err = ts.ParseGlob(filepath.Join(dir, "*.partial.tmpl"))
		if err != nil {
			return nil, err
		}

		cache[name] = ts
	}

	return cache, nil
}