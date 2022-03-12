package main

import (
	"flag"
	"fmt"
	"net/http"
)

type Config struct {
	Port string
	Addr string
}

func main() {
	cfg := Config{}

	// Establish dependencies for handlers
	application := &application{
		infoLogger: InfoLogger,
		errorLogger: ErrorLogger,
	}

	// Parse configurations for the application
	flag.StringVar(&cfg.Addr, "server_id", "0.0.0.0", "server address for witch to listen on")
	flag.StringVar(&cfg.Port, "port", "5555", "server port for witch to listen on")
	flag.Parse()

	srv := &http.Server{
		Addr:    fmt.Sprintf("%s:%s", cfg.Addr, cfg.Port),
		Handler: application.routes(),
		ErrorLog: application.errorLogger,
	}

	application.infoLogger.Printf("Starting server %s:%s", cfg.Addr, cfg.Port)
	application.errorLogger.Fatal(srv.ListenAndServe())
}
