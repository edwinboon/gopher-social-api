package main

import (
	"log"
	"net/http"
)

func (app *application) internalServerErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("internal server error: %s path: %s error: %s", r.Method, r.URL.Path, err.Error())
	writeJSONError(w, http.StatusInternalServerError, "the server encountered a problem and could not process your request")
}

func (app *application) badRequestResponse(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("bad request: %s path: %s error: %s", r.Method, r.URL.Path, err.Error())
	writeJSONError(w, http.StatusBadRequest, err.Error())
}

func (app *application) alreadyExistsResponse(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("value already exists: %s path: %s error: %s", r.Method, r.URL.Path, err.Error())
	writeJSONError(w, http.StatusConflict, err.Error())
}

func (app *application) notFoundResponse(w http.ResponseWriter, r *http.Request) {
	log.Printf("not found: %s path: %s", r.Method, r.URL.Path)
	writeJSONError(w, http.StatusNotFound, "the requested resource could not be found")
}
