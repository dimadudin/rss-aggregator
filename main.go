package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	port := os.Getenv("PORT")
	mux := http.NewServeMux()
	mux.HandleFunc("GET /v1/readiness", handleReady)
	mux.HandleFunc("GET /v1/error", handleError)
	corsMux := MwAddCors(mux)
	server := http.Server{Addr: ":" + port, Handler: corsMux}
	log.Fatal(server.ListenAndServe())
}
