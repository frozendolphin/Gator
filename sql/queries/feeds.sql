-- name: CreateFeed :one
INSERT INTO feeds (id, name, url, user_id)
VALUES (
    $1,
    $2,
    $3,
    $4
)
RETURNING *;

-- name: GetFeeds :many
SELECT feeds.name AS feed_name, feeds.url AS feed_url, users.name AS user_name
FROM feeds
INNER JOIN users
ON feeds.user_id = users.id;

-- name: GetFeedByUrl :one
SELECT * FROM feeds
WHERE url = $1
LIMIT 1;