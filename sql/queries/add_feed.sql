-- name: AddFeed :one
INSERT INTO feeds (
    id,
    created_at,
    updated_at,
    name,
    url
) VALUES (
    $1,
    $2,
    $3,
    $4,
    $5
)
RETURNING *;


