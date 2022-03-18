package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"text/template"

	"github.com/esirk/snippet_box/pkg/models"
)

type PageData struct {
	Title  string
	Header string
}

var fileServer = http.FileServer(http.Dir("./ui/static"))

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.clientError(w, http.StatusNotFound, nil)
		return
	}
	data := PageData{
		Title:  "Learning GO",
		Header: "Learning GO Further",
	}
	files := []string{
		"ui/html/home.page.tmpl",
		"ui/html/base.layout.tmpl",
		"ui/html/footer.partial.tmpl",
	}
	tmpl, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, err)
		return
	}
	if err = tmpl.Execute(w, data); err != nil {
		app.serverError(w, err)
		return
	}
}

func (app *application) showSnippet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		app.clientError(w, http.StatusBadRequest, &models.ClientError{
			Message: "ID must be a positive integer",
		})
		return
	}
	files := []string{
		"ui/html/show.page.tmpl",
		"ui/html/base.layout.tmpl",
		"ui/html/footer.partial.tmpl",
	}
	tmpl, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, err)
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
	err = tmpl.Execute(w, snippet)
	if err != nil {
		app.serverError(w, err)
	}
}

func (app *application)showAllSnippets(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"ui/html/show.page.tmpl",
		"ui/html/base.layout.tmpl",
		"ui/html/footer.partial.tmpl",
	}

	tmpl, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, err)
		return
	}
	snippets, _ := app.snippets.Latest()
	err = tmpl.Execute(w, snippets)
	if err != nil {
		app.serverError(w, err)
	}
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
