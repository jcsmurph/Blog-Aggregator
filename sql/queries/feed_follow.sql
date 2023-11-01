-- name: CreateFeedFollow :one
INSERT INTO rss_feed_follow (
    id,
    created_at,
    updated_at,
    feed_id,
    user_id
    ) VALUES ( $1, $2, $3, $4, $5 )
RETURNING *;

-- name: GetFeedFollowByUser :many
SELECT * FROM rss_feed_follow WHERE user_id = $1;

-- name: DeleteFeedFollow :exec
DELETE FROM rss_feed_follow WHERE feed_id = $1 AND user_id = $2;
