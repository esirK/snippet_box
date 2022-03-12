package main

import "log"

type loggers struct {
	infoLogger *log.Logger
	errorLogger *log.Logger
}

type application struct {
	loggers loggers
}

