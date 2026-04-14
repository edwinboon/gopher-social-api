-- +goose Up
CREATE TABLE IF NOT EXISTS posts (
	id BIGSERIAL PRIMARY KEY,
	title text NOT NULL,
	content text NOT NULL,
	user_id bigint NOT NULL,
	created_at timestamp(0) with time zone NOT NULL DEFAULT now(),
	CONSTRAINT fk_user_id FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE IF EXISTS posts;
