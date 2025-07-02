-- +goose up
CREATE TABLE users (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP not null,
    updated_AT TIMESTAMP not null,
    name TEXT not null 
);

-- +goose down
DROP TABLE users;
