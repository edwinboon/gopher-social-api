-- +goose Up
CREATE TABLE IF NOT EXISTS followers (
	user_id BIGINT NOT NULL,
	follower_id BIGINT NOT NULL,
	created_at timestamp(0) with time zone NOT NULL DEFAULT now(),
	-- composite primary key to prevent duplicate entries
	PRIMARY KEY (user_id, follower_id),
	-- foreign key constraints
	FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
	FOREIGN KEY (follower_id) REFERENCES users(id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE IF EXISTS followers;
