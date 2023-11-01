-- +goose Up
ALTER TABLE rss_feed ADD COLUMN last_fetched TIMESTAMP;

-- +goose Down
ALTER TABLE rss_feed DROP COLUMN last_fetched;
