package main

import (
	"log"
	"net/http"
	"time"
)

type config struct {
	addr string
}

type application struct {
	config config
}

func (app *application) mount() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /v1/health", app.healthCheckHandler)

	return mux
}

func (app *application) serve(mux *http.ServeMux) error {
	srv := &http.Server{
		Addr:         app.config.addr,
		Handler:      mux,
		WriteTimeout: 30 * time.Second,
		ReadTimeout:  10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	log.Printf("Starting server on %s", app.config.addr)

	return srv.ListenAndServe()
}
