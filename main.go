package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/dimadudin/rss-aggregator/internal/database"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type config struct {
	DB *database.Queries
}

func main() {
	godotenv.Load()
	port := os.Getenv("PORT")
	dbURL := os.Getenv("CONN")
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		panic(err.Error())
	}
	dbQueries := database.New(db)
	cfg := config{DB: dbQueries}

	mux := http.NewServeMux()
	mux.HandleFunc("POST /v1/users", handleCreateUser)
	mux.HandleFunc("GET /v1/readiness", handleReady)
	mux.HandleFunc("GET /v1/error", handleError)
	corsMux := MwAddCors(mux)
	server := http.Server{Addr: ":" + port, Handler: corsMux}
	log.Fatal(server.ListenAndServe())
}
