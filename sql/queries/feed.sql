-- name: CreateFeed :one
INSERT INTO rss_feed (
    id,
    created_at,
    updated_at,
    name,
    url,
    user_id,
    last_fetched
    ) VALUES ( $1, $2, $3, $4, $5, $6, $7 )
RETURNING *;

-- name: GetAllRssFeeds :many
SELECT * FROM rss_feed;

-- name: GetNextFeedsToFetch :many
SELECT * FROM rss_feed ORDER BY last_fetched;
