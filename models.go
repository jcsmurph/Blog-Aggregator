package main

import (
	"time"

	"github.com/google/uuid"
	"github.com/jscmurph/blog_aggregator/internal/database"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	APIKey    string    `json:"api_key"`
}

func databaseUserToUser(user database.User) User {
	return User{
		ID:        user.ID,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Name:      user.Name,
		APIKey:    user.ApiKey,
	}
}

type RssFeed struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	Url       string    `json:"url"`
	UserId   uuid.UUID `json:"user_id"`
}

func databaseFeedToFeed(rssFeed database.RssFeed) RssFeed {
	return RssFeed{
		ID:        rssFeed.ID,
		CreatedAt: rssFeed.CreatedAt,
		UpdatedAt: rssFeed.UpdatedAt,
		Name:      rssFeed.Name,
		Url:       rssFeed.Url,
		UserId:   rssFeed.UserID,
	}
}
func databaseFeedsToFeeds(rssFeeds []database.RssFeed) []RssFeed {
    result := make([]RssFeed, len(rssFeeds))

	for i, feed := range rssFeeds {
        result[i] = databaseFeedToFeed(feed)
    }

    return result
}
