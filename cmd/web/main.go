package main

import (
	"database/sql"
	"flag"
	"fmt"
	"net/http"

	"github.com/esirk/snippet_box/pkg/models/mysql"
	_ "github.com/go-sql-driver/mysql"
)

type Config struct {
	Port string
	Addr string
}

func main() {
	cfg := Config{}

	db, err := openDB("snippet:snippet@/snippetbox?parseTime=true")

	if err != nil {
		panic(err)
	}
	defer db.Close()
	// Establish dependencies for handlers
	loggers := loggers{
		infoLogger:  InfoLogger,
		errorLogger: ErrorLogger,
	}

	snippetModel := mysql.SnippetModel{
		DB: db,
	}

	application := &application{
		loggers:  &loggers,
		snippets: &snippetModel,
	}

	// Parse configurations for the application
	flag.StringVar(&cfg.Addr, "server_id", "0.0.0.0", "server address for witch to listen on")
	flag.StringVar(&cfg.Port, "port", "5555", "server port for witch to listen on")
	flag.Parse()

	srv := &http.Server{
		Addr:     fmt.Sprintf("%s:%s", cfg.Addr, cfg.Port),
		Handler:  application.routes(),
		ErrorLog: application.loggers.errorLogger,
	}

	application.loggers.infoLogger.Printf("Starting server %s:%s", cfg.Addr, cfg.Port)
	application.loggers.errorLogger.Fatal(srv.ListenAndServe())
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
