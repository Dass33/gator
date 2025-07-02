-- name: GetFeeds :many
select feeds.* from feeds
order BY last_fetched_at asc nulls first;


