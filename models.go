package main

import (
	"database/sql"
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
	ID          uuid.UUID `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Name        string    `json:"name"`
	Url         string    `json:"url"`
	UserId      uuid.UUID `json:"user_id"`
	LastFetched time.Time `json:"last_fetched"`
}

func databaseFeedToFeed(rssFeed database.RssFeed) RssFeed {
	return RssFeed{
		ID:          rssFeed.ID,
		CreatedAt:   rssFeed.CreatedAt,
		UpdatedAt:   rssFeed.UpdatedAt,
		Name:        rssFeed.Name,
		Url:         rssFeed.Url,
		UserId:      rssFeed.UserID,
		LastFetched: *nullTimeToTimePtr(rssFeed.LastFetched),
	}
}

func databaseFeedsToFeeds(rssFeeds []database.RssFeed) []RssFeed {
	result := make([]RssFeed, len(rssFeeds))

	for i, feed := range rssFeeds {
		result[i] = databaseFeedToFeed(feed)
	}

	return result
}

type RssFeedFollow struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	FeedId    uuid.UUID `json:"feed_id"`
	UserId    uuid.UUID `json:"user_id"`
}

func databaseFeedFollowToFeedFollow(rssFeedFollow database.RssFeedFollow) RssFeedFollow {
	return RssFeedFollow{
		ID:        rssFeedFollow.ID,
		CreatedAt: rssFeedFollow.CreatedAt,
		UpdatedAt: rssFeedFollow.UpdatedAt,
		FeedId:    rssFeedFollow.FeedID,
		UserId:    rssFeedFollow.UserID,
	}
}

func databaseFeedFollowsToFeedFollows(feedFollows []database.RssFeedFollow) []RssFeedFollow {
	result := make([]RssFeedFollow, len(feedFollows))
	for i, feedFollow := range feedFollows {
		result[i] = databaseFeedFollowToFeedFollow(feedFollow)
	}
	return result
}

type Post struct {
	ID           uuid.UUID `json:"id"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	Title        string    `json:"title"`
	Url          string    `json:"url"`
	Description  string    `json:"description"`
	Published_At time.Time `json:"published_at"`
	FeedID       uuid.UUID `json:"feed_id"`
}

func databasePostToPost(post database.Post) Post {
	return Post{
		ID:           post.ID,
		CreatedAt:    post.CreatedAt,
		UpdatedAt:    post.UpdatedAt,
		Title:        post.Title,
		Url:          post.Url,
		Description:  *nullStringToStringPtr(post.Description),
		Published_At: *nullTimeToTimePtr(post.PublishedAt),
		FeedID:       post.FeedID,
	}
}

func databasePostsToPosts(posts []database.Post) []Post {
	result := make([]Post, len(posts))
	for i, post := range posts {
		result[i] = databasePostToPost(post)
	}
	return result
}

func nullStringToStringPtr(s sql.NullString) *string {
	if s.Valid {
		return &s.String
	}
	return nil
}

func nullTimeToTimePtr(t sql.NullTime) *time.Time {
	if t.Valid {
		return &t.Time
	}
	return nil
}
