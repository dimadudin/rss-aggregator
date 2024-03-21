package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
)

func handleReady(w http.ResponseWriter, r *http.Request) {
	response := struct {
		Status string `json:"status"`
	}{Status: "ok"}
	RespondWithJSON(w, http.StatusOK, response)
}

func handleError(w http.ResponseWriter, r *http.Request) {
	RespondWithError(w, http.StatusInternalServerError, "Internal Server Error")
}

func (cfg *config) handleCreateUser(w http.ResponseWriter, r *http.Request) {
	request := struct {
		Name string `json:"name"`
	}{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&request)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	response := struct {
		Id        uuid.UUID `json:"id"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"update_at"`
		Name      string    `json:"name"`
	}{
		Id:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      request.Name,
	}
	cfg.DB.CreateUser()
	RespondWithError(w, http.StatusInternalServerError, "Internal Server Error")
}
