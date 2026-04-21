package main

import "net/http"

func (app *application) healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"status":  "available",
		"env":     app.config.env,
		"version": app.config.version,
	}

	// use just writeJSON instead of jsonResponse to avoid wrapping the response in a "data" envelope
	if err := writeJSON(w, http.StatusOK, data); err != nil {
		app.internalServerErrorResponse(w, r, err)
	}
}
