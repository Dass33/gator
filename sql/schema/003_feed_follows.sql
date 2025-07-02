-- +goose up
CREATE TABLE feed_follows (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP not null,
    updated_at TIMESTAMP not null,
    user_id UUID  not null,
    feed_id UUID  not null,
    CONSTRAINT fk_users
    FOREIGN KEY (user_id)
        REFERENCES users(id)
        ON DELETE CASCADE,

    CONSTRAINT fk_feed
    FOREIGN KEY (feed_id)
        REFERENCES feeds(id)
        ON DELETE CASCADE,

    unique (user_id, feed_id)
);
ALTER TABLE feeds 
DROP user_id;

-- +goose down
DROP TABLE feeds_follows;

ALTER TABLE feeds
ADD COLUMN user_id UUID;

ALTER TABLE feeds
ADD CONSTRAINT fk_users
FOREIGN KEY (user_id)
    REFERENCES users(id)
    ON DELETE CASCADE;
