package main

import (
	"net/http"

	"github.com/go-chi/chi"
)

func (app *application) routes() http.Handler {
	// mux := http.NewServeMux()
	router := chi.NewRouter()
	router.Get("/", app.home)
	router.Get("/snippets", app.home)
	router.Post("/snippets/create", app.createSnippet)
	router.Get("/snippets/create", app.createSnippetForm)
	router.Get("/snippets/{id}", app.showSnippet)
	router.Delete("/snippets/{id}", app.deleteSnippet)
	router.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return app.recoverPanic(app.logRequest(secureHeaders(router)))
}
