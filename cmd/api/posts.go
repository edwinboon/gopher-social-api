package main

import (
	"net/http"

	"github.com/edwinboon/gopher-social-api/internal/store"
)

type CreatePostPayload struct {
	Title   string   `json:"title"`
	Content string   `json:"content"`
	Tags    []string `json:"tags"`
}

func (app *application) createPostHandler(w http.ResponseWriter, r *http.Request) {
	var payload CreatePostPayload

	if err := ReadJSON(w, r, &payload); err != nil {
		WriteJSONError(w, http.StatusBadRequest, err.Error())
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
		WriteJSONError(w, http.StatusInternalServerError, err.Error())
		return
	}

	if err := WriteJSON(w, http.StatusCreated, post); err != nil {
		WriteJSONError(w, http.StatusInternalServerError, err.Error())
		return
	}
}
