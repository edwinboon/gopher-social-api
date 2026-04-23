-- +goose Up
ALTER TABLE posts ADD COLUMN version INTEGER NOT NULL DEFAULT 0;

-- +goose Down
ALTER TABLE posts DROP COLUMN version;
