package main

import "github.com/esirk/snippet_box/pkg/models"

type SnippetData struct {
	Snippet *models.Snippet
}

type SnippetsData struct {
	Snippets []*models.Snippet
}
