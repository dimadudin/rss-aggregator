package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/dimadudin/rss-aggregator/internal/database"
	"github.com/google/uuid"
)

func handleReady(w http.ResponseWriter, r *http.Request) {
	response := struct {
		Status string `json:"status"`
	}{Status: "ok"}
	respondWithJSON(w, http.StatusOK, response)
}

func handleError(w http.ResponseWriter, r *http.Request) {
	respondWithError(w, http.StatusInternalServerError, "Internal Server Error")
}

func (cfg *config) handleCreateUser(w http.ResponseWriter, r *http.Request) {
	rParams := struct {
		Name string `json:"name"`
	}{}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&rParams)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	user, err := cfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      rParams.Name,
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, databaseUserToUser(user))
}

func (cfg *config) handleGetUser(w http.ResponseWriter, r *http.Request, user database.User) {
	respondWithJSON(w, http.StatusOK, databaseUserToUser(user))
}

func (cfg *config) handleCreateFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	rParams := struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	}{}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&rParams)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
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
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, databaseFeedToFeed(feed))
}
