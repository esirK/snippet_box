package main

import (
	"html/template"
	"log"

	"github.com/esirk/snippet_box/pkg/models/mysql"
)

type loggers struct {
	infoLogger  *log.Logger
	errorLogger *log.Logger
	netLogger   *log.Logger
}

type application struct {
	loggers  *loggers
	snippets *mysql.SnippetModel
	templateCache map[string]*template.Template
}
