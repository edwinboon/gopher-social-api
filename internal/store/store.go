package store

import (
	"context"
	"database/sql"
	"errors"
)

var ErrNotFound = errors.New("resource not found")

type Store struct {
	Posts interface {
		GetByID(context.Context, int64) (*Post, error)
		Create(context.Context, *Post) error
		Delete(context.Context, int64) error
	}
	Users interface {
		Create(context.Context, *User) error
	}
	Comments interface {
		GetByPostID(context.Context, int64) ([]CommentWithUser, error)
	}
}

func NewStore(db *sql.DB) Store {
	return Store{
		Posts:    &PostStore{db},
		Users:    &UserStore{db},
		Comments: &CommentStore{db},
	}
}
