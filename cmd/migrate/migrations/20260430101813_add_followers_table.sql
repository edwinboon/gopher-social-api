-- +goose Up
CREATE TABLE IF NOT EXISTS followers (
	user_id BIGINT NOT NULL,
	follower_id BIGINT NOT NULL,
	created_at timestamp(0) with time zone NOT NULL DEFAULT now(),

	PRIMARY KEY (user_id, follower_id),

	CONSTRAINT followers_user_id_fkey FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
	CONSTRAINT followers_follower_id_fkey FOREIGN KEY (follower_id) REFERENCES users(id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE IF EXISTS followers;
