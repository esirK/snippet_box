package main

import (
	"log"

	"github.com/esirk/snippet_box/pkg/models/mysql"
)

type loggers struct {
	infoLogger  *log.Logger
	errorLogger *log.Logger
}

type application struct {
	loggers  *loggers
	snippets *mysql.SnippetModel
}
