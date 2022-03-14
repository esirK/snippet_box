package main

import (
	"database/sql"
	"log"
)

type loggers struct {
	infoLogger  *log.Logger
	errorLogger *log.Logger
}

type application struct {
	loggers *loggers
	db      *sql.DB
}
