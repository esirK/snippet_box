package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/esirk/snippet_box/pkg/models"
	"github.com/go-chi/chi"
)

var fileServer = http.FileServer(http.Dir("./ui/static"))

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	snippets, _ := app.snippets.Latest()
	data := &templateData{
		Snippets: snippets,
	}
	app.render(w, "home.page.tmpl", data)
}

func (app *application) showSnippet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil || id < 1 {
		app.clientError(w, http.StatusBadRequest, &models.ClientError{
			Message: "ID must be a positive integer",
		})
		return
	}

	snippet, err := app.snippets.Get(id)
	if err != nil {
		switch err.(type) {
		case *models.ErrNoRecord:
			app.clientError(w, http.StatusNotFound, &models.ClientError{
				Message: "Snippet not found",
			})
			return
		default:
			app.serverError(w, err)
		}
		return
	}
	data := &templateData{
		Snippet: snippet,
	}
	app.render(w, "snippet.details.page.tmpl", data)
}

func (app *application) createSnippet(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		// w.WriteHeader(http.StatusMethodNotAllowed)
		// w.Write([]byte("Method not allowed"))
		app.clientError(w, http.StatusMethodNotAllowed, nil)
		return
	}
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		app.serverError(w, err)
		return
	}
	var snippet models.Snippet
	if err := json.Unmarshal(data, &snippet); err != nil {
		app.serverError(w, err)
		return
	}
	id, err := app.snippets.Insert(&snippet)
	if err != nil {
		app.serverError(w, err)
		return
	}
	// redirect to new snippet
	http.Redirect(w, r, fmt.Sprintf("/snippet?id=%d", id), http.StatusSeeOther)
}
