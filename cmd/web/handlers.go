package main

import (
	"fmt"
	"net/http"
	"strconv"
	"text/template"
)

type PageData struct {
	Title  string
	Header string
}

var fileServer = http.FileServer(http.Dir("./ui/static"))

func (app *application)home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.clientError(w, http.StatusNotFound)
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

func (app *application)showSnippet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	fmt.Printf("id: %d, err: %v\n", id, err)

	if id < 1 && err != nil {
		w.Write([]byte("Showing all snippets"))
		return
	}
	if err != nil || id < 1 {
		app.clientError(w, http.StatusBadRequest)
		return
	} else {
		fmt.Fprintf(w, "Showing snippet: %d", id)
	}
}

func (app *application)createSnippet(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		// w.WriteHeader(http.StatusMethodNotAllowed)
		// w.Write([]byte("Method not allowed"))
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}
	w.Write([]byte("Create a new snippet"))
}
