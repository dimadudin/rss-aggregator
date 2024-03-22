package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/dimadudin/rss-aggregator/internal/database"
	"github.com/google/uuid"
)

func (cfg *config) handleFollowFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	rParams := struct {
		FeedID uuid.UUID `json:"feed_id"`
	}{}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&rParams)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError,
			fmt.Sprintf("error decoding JSON: %s", err.Error()))
		return
	}

	follow, err := cfg.DB.CreateFollow(r.Context(), database.CreateFollowParams{
		ID:        uuid.New(),
		UserID:    user.ID,
		FeedID:    rParams.FeedID,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError,
			fmt.Sprintf("error following feed: %s", err.Error()))
		return
	}

	respondWithJSON(w, http.StatusCreated, databaseFollowToFollow(follow))
}

func (cfg *config) handleUnfollowFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	feedID, err := uuid.Parse(r.PathValue("followID"))
	if err != nil {
		respondWithError(w, http.StatusInternalServerError,
			fmt.Sprintf("error parsing id: %s", err.Error()))
		return
	}

	follow, err := cfg.DB.DeleteFollow(r.Context(), feedID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError,
			fmt.Sprintf("error unfollowing feed: %s", err.Error()))
		return
	}

	respondWithJSON(w, http.StatusOK, databaseFollowToFollow(follow))
}
