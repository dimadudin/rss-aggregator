package main

import (
	"context"
	"database/sql"
	"log"
	"sync"
	"time"

	"github.com/dimadudin/rss-aggregator/internal/database"
	"github.com/google/uuid"
)

func startScraping(db *database.Queries, requestCap int32, requestInterval time.Duration) {
	log.Printf("Initiating scrape worker on %v goroutines every %v", requestCap, requestInterval)

	ticker := time.NewTicker(requestInterval)
	for ; ; <-ticker.C {
		feeds, err := db.GetNextFeedsToFetch(context.Background(), requestCap)
		if err != nil {
			log.Println("error fetching feeds:", err.Error())
			continue
		}

		wg := &sync.WaitGroup{}
		for _, feed := range feeds {
			wg.Add(1)
			go scrapeFeed(db, wg, feed)
		}
		wg.Wait()
	}
}

func scrapeFeed(db *database.Queries, wg *sync.WaitGroup, feed database.Feed) {
	defer wg.Done()

	_, err := db.MarkFeedFetched(context.Background(), feed.ID)
	if err != nil {
		log.Println("error marking feed as fetched:", err.Error())
		return
	}

	rssFeed, err := urlToFeed(feed.Url)
	if err != nil {
		log.Println("error fetching feed:", err.Error())
		return
	}

	for _, item := range rssFeed.Channel.Item {
		desc := sql.NullString{}
		if item.Description != "" {
			desc.String = item.Description
			desc.Valid = true
		}

		pubAt, err := time.Parse(time.RFC1123Z, item.PubDate)
		if err != nil {
			log.Printf("error parsing publish date %v: %s", item.PubDate, err.Error())
		}

		_, err = db.CreatePost(context.Background(), database.CreatePostParams{
			ID:          uuid.New(),
			CreatedAt:   time.Now().UTC(),
			UpdatedAt:   time.Now().UTC(),
			Title:       item.Title,
			Url:         item.Link,
			Description: desc,
			PublishedAt: pubAt,
			FeedID:      feed.ID,
		})
		if err != nil {
			log.Println("error creating post:", err.Error())
			return
		}
	}
	log.Printf("Feed %s collected, %v posts found\n", feed.Name, len(rssFeed.Channel.Item))
}
