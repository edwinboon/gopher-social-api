package main

import (
	"net/http"
	"strconv"

	"github.com/edwinboon/gopher-social-api/internal/store"
	"github.com/go-chi/chi/v5"
)

func (app *application) getUserHandler(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.ParseInt(chi.URLParam(r, "userId"), 10, 64)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}
	ctx := r.Context()

	user, err := app.store.Users.GetByID(ctx, userID)
	if err != nil {
		switch err {
		case store.ErrNotFound:
			app.notFoundResponse(w, r)
		default:
			app.internalServerErrorResponse(w, r, err)
		}
		return
	}

	if err := app.jsonResponse(w, http.StatusOK, user); err != nil {
		app.internalServerErrorResponse(w, r, err)
		return
	}
}
