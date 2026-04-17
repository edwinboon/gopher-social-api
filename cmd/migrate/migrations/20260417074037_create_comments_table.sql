-- +goose Up
CREATE TABLE IF NOT EXISTS comments(
	id bigserial PRIMARY KEY,
	post_id bigint NOT NULL,
	user_id bigint NOT NULL,
	content text NOT NULL,
	created_at timestamp(0) with time zone NOT NULL DEFAULT now(),
	CONSTRAINT comments_post_id_fkey FOREIGN KEY (post_id) REFERENCES posts(id) ON DELETE CASCADE,
	CONSTRAINT comments_user_id_fkey FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

CREATE INDEX IF NOT EXISTS idx_comments_post_id ON comments(post_id);

-- +goose Down
DROP INDEX IF EXISTS idx_comments_post_id;
DROP TABLE IF EXISTS comments;
