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
	router.Post("/snippets", app.createSnippet)
	router.Get("/snippets/{id}", app.showSnippet)
	router.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return app.recoverPanic(app.logRequest(secureHeaders(router)))
}
