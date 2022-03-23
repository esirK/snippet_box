package main

import (
	"bytes"
	"fmt"
	"net/http"
	"runtime/debug"
	"time"

	"github.com/esirk/snippet_box/pkg/models"
)

func (app *application) serverError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.loggers.errorLogger.Output(2, trace)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func (app *application) clientError(w http.ResponseWriter, status int, err *models.ClientError) {
	if err != nil {
		http.Error(w, err.Error(), status)
		return
	}
	http.Error(w, http.StatusText(status), status)
}

func (app *application) addDefaultYear(td *templateData)*templateData {
	if td == nil {
		td = &templateData{}
	}
	td.CurrentYear = time.Now().Year()
	return td
}


func (app *application) render(w http.ResponseWriter, name string, td *templateData){
	// retrieve appropriate template from cache
	ts, ok := app.templateCache[name]
	if !ok {
		app.serverError(w, fmt.Errorf("template %s not found", name))
		return
	}
	buf := new(bytes.Buffer)

	err := ts.Execute(buf, app.addDefaultYear(td))
	if err != nil {
		app.serverError(w, err)
	}
	buf.WriteTo(w)
}
