-- +goose Up
ALTER TABLE posts ALTER COLUMN title TYPE varchar(100);

-- +goose Down
ALTER TABLE posts ALTER COLUMN title TYPE text;
