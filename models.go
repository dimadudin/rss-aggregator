package main

import (
	"time"

	"github.com/dimadudin/rss-aggregator/internal/database"
	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	ApiKey    string    `json:"api_key"`
}

func databaseUserToUser(dbUser database.User) User {
	return User{
		ID:        dbUser.ID,
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: dbUser.UpdatedAt,
		Name:      dbUser.Name,
		ApiKey:    dbUser.ApiKey,
	}
}

type Feed struct {
	ID            uuid.UUID  `json:"id"`
	CreatedAt     time.Time  `json:"created_at"`
	UpdatedAt     time.Time  `json:"updated_at"`
	Name          string     `json:"name"`
	Url           string     `json:"url"`
	UserID        uuid.UUID  `json:"user_id"`
	LastFetchedAt *time.Time `json:"last_fetched_at"`
}

func databaseFeedToFeed(dbFeed database.Feed) Feed {
	var lastFetchedAtTime *time.Time
	if dbFeed.LastFetchedAt.Valid {
		lastFetchedAtTime = &dbFeed.LastFetchedAt.Time
	} else {
		lastFetchedAtTime = nil
	}

	return Feed{
		ID:            dbFeed.ID,
		CreatedAt:     dbFeed.CreatedAt,
		UpdatedAt:     dbFeed.UpdatedAt,
		Name:          dbFeed.Name,
		Url:           dbFeed.Url,
		UserID:        dbFeed.UserID,
		LastFetchedAt: lastFetchedAtTime,
	}
}

func databaseFeedsToFeeds(dbFeeds []database.Feed) []Feed {
	feeds := make([]Feed, 0, len(dbFeeds))
	for _, dbFeed := range dbFeeds {
		feeds = append(feeds, databaseFeedToFeed(dbFeed))
	}
	return feeds
}

type Follow struct {
	ID        uuid.UUID `json:"id"`
	UserID    uuid.UUID `json:"user_id"`
	FeedID    uuid.UUID `json:"feed_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func databaseFollowToFollow(dbFollow database.Follow) Follow {
	return Follow{
		ID:        dbFollow.ID,
		UserID:    dbFollow.UserID,
		FeedID:    dbFollow.FeedID,
		CreatedAt: dbFollow.CreatedAt,
		UpdatedAt: dbFollow.UpdatedAt,
	}
}

func databaseFollowsToFollows(dbFollows []database.Follow) []Follow {
	follows := make([]Follow, 0, len(dbFollows))
	for _, dbFeed := range dbFollows {
		follows = append(follows, databaseFollowToFollow(dbFeed))
	}
	return follows
}

//
// type Post struct {
// 	ID          uuid.UUID `json:"id"`
// 	CreatedAt   time.Time `json:"created_at"`
// 	UpdatedAt   time.Time `json:"updated_at"`
// 	Title       string    `json:"title"`
// 	Url         string    `json:"url"`
// 	Description *string   `json:"description"`
// 	PublishedAt time.Time `json:"published_at"`
// 	FeedID      uuid.UUID `json:"feed_id"`
// }
//
// func databasePostToPost(dbPost database.Post) Post {
// 	return Post{
// 		ID:          dbPost.ID,
// 		CreatedAt:   dbPost.CreatedAt,
// 		UpdatedAt:   dbPost.UpdatedAt,
// 		Title:       dbPost.Title,
// 		Url:         dbPost.Url,
// 		Description: &dbPost.Description.String,
// 		PublishedAt: dbPost.PublishedAt,
// 		FeedID:      dbPost.FeedID,
// 	}
// }
