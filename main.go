package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"time"

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

	conn, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal(err.Error())
	}

	dbQueries := database.New(conn)
	cfg := config{DB: dbQueries}

	go startScraping(dbQueries, 10, time.Minute)

	mux := http.NewServeMux()
	mux.HandleFunc("GET /v1/readiness", handleReady)
	mux.HandleFunc("GET /v1/error", handleError)

	mux.HandleFunc("POST /v1/users", cfg.handleCreateUser)
	mux.Handle("GET /v1/users", cfg.mwAuth(cfg.handleGetUser))

	mux.Handle("POST /v1/feeds", cfg.mwAuth(cfg.handleCreateFeed))
	mux.HandleFunc("GET /v1/feeds", cfg.handleGetFeeds)

	mux.Handle("GET /v1/feed_follows", cfg.mwAuth(cfg.handleGetFollows))
	mux.Handle("POST /v1/feed_follows", cfg.mwAuth(cfg.handleFollowFeed))
	mux.Handle("DELETE /v1/feed_follows/{followID}", cfg.mwAuth(cfg.handleUnfollowFeed))

	mux.Handle("GET /v1/posts", cfg.mwAuth(cfg.handleGetPosts))

	corsMux := mwAddCors(mux)

	server := http.Server{Addr: ":" + port, Handler: corsMux}
	log.Fatal(server.ListenAndServe())
}
