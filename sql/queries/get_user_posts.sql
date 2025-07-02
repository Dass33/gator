-- name: UserPosts :many
select * from posts p
join feed_follows ff
on p.feed_id = ff.feed_id and  ff.user_id = $1
order by published_at desc
limit $2;
