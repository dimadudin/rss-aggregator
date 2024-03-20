package main

import "net/http"

func handleReady(w http.ResponseWriter, r *http.Request) {
	response := struct {
		Status string `json:"status"`
	}{Status: "ok"}
	RespondWithJSON(w, http.StatusOK, response)
}

func handleError(w http.ResponseWriter, r *http.Request) {
	RespondWithError(w, http.StatusInternalServerError, "Internal Server Error")
}
