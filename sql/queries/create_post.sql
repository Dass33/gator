-- name: CreatePost :one
insert into posts (
    id,
    created_at,
    updated_at,
    published_at,
    url,
    title,
    description,
    feed_id
) VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6,
    $7,
    $8
)
RETURNING *;

