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

func (app *application) showSnippet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	fmt.Printf("id: %d, err: %v\n", id, err)

	if id < 1 && err != nil {
		snippets, _ := app.snippets.Latest()
		var dat string
		for _, s := range snippets {
			dat += fmt.Sprintf("%s\n", s.Content)
		}
		w.Write([]byte(dat))
		return
	}
	if err != nil || id < 1 {
		app.clientError(w, http.StatusBadRequest)
		return
	} else {
		snippet, _ := app.snippets.Get(id)
		fmt.Fprintf(w, "%v", snippet)
	}
}

func (app *application) createSnippet(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		// w.WriteHeader(http.StatusMethodNotAllowed)
		// w.Write([]byte("Method not allowed"))
		app.clientError(w, http.StatusMethodNotAllowed)
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
