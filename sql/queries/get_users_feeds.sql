-- name: GetUsersFeeds :many
select * from feeds
where feeds.id in
(
    select ff.feed_id from feed_follows ff
    where ff.user_id = $1
);
