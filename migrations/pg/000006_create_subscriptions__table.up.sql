CREATE TABLE IF NOT EXISTS subscriptions (
    user_id INT REFERENCES users(id) ON DELETE CASCADE NOT NULL,
    channel_id INT REFERENCES channels(id) ON DELETE CASCADE NOT NULL,
    role VARCHAR(10) NOT NULL CHECK(role IN ('USER', 'ADMIN', 'MODERATOR')),
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP NOT NULL DEFAULT now(),
    PRIMARY KEY(user_id, channel_id)
)