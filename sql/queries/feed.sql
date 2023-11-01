-- name: CreateFeed :one
INSERT INTO rss_feed (
    id,
    created_at,
    updated_at,
    name,
    url,
    user_id
    ) VALUES ( $1, $2, $3, $4, $5, $6 )
RETURNING *;

-- name: GetAllRssFeeds :many
SELECT * FROM rss_feed;
