package store

import (
	"context"
	"database/sql"
	"errors"
	"time"
)

var (
	ErrNotFound      = errors.New("resource not found")
	ErrAlreadyExists = errors.New("resource already exists")
)

const QueryTimeoutDuration = 5 * time.Second

type Store struct {
	Posts interface {
		GetByID(context.Context, int64) (*Post, error)
		Create(context.Context, *Post) error
		Delete(context.Context, int64) error
		Patch(context.Context, *Post) (*Post, error)
	}
	Users interface {
		Create(context.Context, *User) error
		GetByID(context.Context, int64) (*User, error)
	}
	Comments interface {
		GetByPostID(context.Context, int64) ([]CommentWithUser, error)
		Create(context.Context, *Comment) error
	}
	Followers interface {
		Follow(ctx context.Context, userID, followerID int64) error
		Unfollow(ctx context.Context, userID, followerID int64) error
	}
}

func NewStore(db *sql.DB) Store {
	return Store{
		Posts:     &PostStore{db},
		Users:     &UserStore{db},
		Comments:  &CommentStore{db},
		Followers: &FollowStore{db},
	}
}
