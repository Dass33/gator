-- +goose up
CREATE TABLE feeds (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP not null,
    updated_at TIMESTAMP not null,
    name TEXT not null, 
    url TEXT  not null unique,
    user_id UUID  not null,
    CONSTRAINT fk_users
    FOREIGN KEY (user_id)
        REFERENCES users(id)
        ON DELETE CASCADE
);

-- +goose down
DROP TABLE feeds;
