-- name: GetPost :one
select * from Posts
where published_at = $1
and url = $2
and title = $3;
