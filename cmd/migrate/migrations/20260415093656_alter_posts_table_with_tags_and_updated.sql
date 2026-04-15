-- +goose Up
ALTER TABLE posts
ADD COLUMN tags text[] NOT NULL DEFAULT '{}',
ADD COLUMN updated_at timestamp(0) with time zone NOT NULL DEFAULT now();

-- +goose Down
ALTER TABLE posts
DROP COLUMN IF EXISTS tags,
DROP COLUMN IF EXISTS updated_at;
