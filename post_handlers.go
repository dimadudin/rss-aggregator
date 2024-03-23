package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/dimadudin/rss-aggregator/internal/database"
)

func (cfg *config) handleGetPosts(w http.ResponseWriter, r *http.Request, user database.User) {
	postLimitStr := r.URL.Query().Get("limit")
	if postLimitStr == "" {
		postLimitStr = "10"
	}

	postLimit, err := strconv.Atoi(postLimitStr)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError,
			fmt.Sprintf("error parsing limit query parameter: %s", err.Error()))
		return
	}

	posts, err := cfg.DB.GetPostsByUser(r.Context(), database.GetPostsByUserParams{
		UserID: user.ID,
		Limit:  int32(postLimit),
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError,
			fmt.Sprintf("error fetching posts for user: %s", err.Error()))
		return
	}

	respondWithJSON(w, http.StatusOK, databasePostsToPosts(posts))
}
