package main

import "net/http"

func (app *application) healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"status":  "available",
		"env":     app.config.env,
		"version": app.config.version,
	}

	if err := app.JsonResponse(w, http.StatusOK, data); err != nil {
		app.internalServerErrorResponse(w, r, err)
	}
}
