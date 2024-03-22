package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/dimadudin/rss-aggregator/internal/database"
	"github.com/google/uuid"
)

func (cfg *config) handleCreateFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	rParams := struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	}{}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&rParams)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError,
			fmt.Sprintf("error decoding JSON: %s", err.Error()))
		return
	}

	feed, err := cfg.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      rParams.Name,
		Url:       rParams.URL,
		UserID:    user.ID,
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError,
			fmt.Sprintf("error creating feed: %s", err.Error()))
		return
	}

	respondWithJSON(w, http.StatusCreated, databaseFeedToFeed(feed))
}

func (cfg *config) handleGetFeeds(w http.ResponseWriter, r *http.Request) {
	feeds, err := cfg.DB.GetFeeds(r.Context())
	if err != nil {
		respondWithError(w, http.StatusInternalServerError,
			fmt.Sprintf("error fetching feeds: %s", err.Error()))
		return
	}

	respondWithJSON(w, http.StatusCreated, databaseFeedsToFeeds(feeds))
}
