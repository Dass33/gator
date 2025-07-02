-- +goose up
create table posts (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP not null,
    updated_at TIMESTAMP not null,
    published_at TIMESTAMP not null,
    url TEXT  not null unique,
    title TEXT not null,
    description TEXT not null, 
    feed_id UUID not null,

    CONSTRAINT fk_feed
    FOREIGN KEY (feed_id)
        REFERENCES feeds(id)
        ON DELETE CASCADE
);

-- +goose down
drop table posts;
