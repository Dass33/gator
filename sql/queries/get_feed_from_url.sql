-- name: GetFeedFromUrl :one
select * from feeds
where url = $1;

