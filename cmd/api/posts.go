package main

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/edwinboon/gopher-social-api/internal/store"
	"github.com/go-chi/chi/v5"
)

type CreatePostPayload struct {
	Title   string   `json:"title"`
	Content string   `json:"content"`
	Tags    []string `json:"tags"`
}

func (app *application) createPostHandler(w http.ResponseWriter, r *http.Request) {
	var payload CreatePostPayload

	if err := ReadJSON(w, r, &payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	post := &store.Post{
		Title:   payload.Title,
		Content: payload.Content,
		Tags:    payload.Tags,
		UserID:  1, // TODO: get user ID from authentication
	}
	ctx := r.Context()
	if err := app.store.Posts.Create(ctx, post); err != nil {
		app.internalServerErrorResponse(w, r, err)
		return
	}

	if err := WriteJSON(w, http.StatusCreated, post); err != nil {
		app.internalServerErrorResponse(w, r, err)
		return
	}
}

func (app *application) getPostHandler(w http.ResponseWriter, r *http.Request) {
	postIDParam := chi.URLParam(r, "postId")

	postID, err := strconv.ParseInt(postIDParam, 10, 64)
	if err != nil {
		app.badRequestResponse(w, r, errors.New("invalid post ID"))
		return
	}

	ctx := r.Context()

	post, err := app.store.Posts.GetByID(ctx, postID)
	if err != nil {
		switch {
		case errors.Is(err, store.ErrNotFound):
			app.notFoundResponse(w, r)
		default:
			app.internalServerErrorResponse(w, r, err)
		}
		return
	}

	if err := WriteJSON(w, http.StatusOK, post); err != nil {
		app.internalServerErrorResponse(w, r, err)
		return
	}
}
