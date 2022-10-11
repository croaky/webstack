package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/url"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
)

func main() {
	// env
	port, ok := os.LookupEnv("PORT")
	if !ok {
		port = "8080"
	}
	dbUrl, ok := os.LookupEnv("DATABASE_URL")
	if !ok {
		dbUrl = "postgres:///webstack_dev"
	}
	primary := os.Getenv("PRIMARY_REGION")
	current := os.Getenv("FLY_REGION")

	// if not in primary region...
	// https://fly.io/docs/getting-started/multi-region-databases/#connect-to-regional-replicas
	if primary != "" && current != "" && primary != current {
		u, err := url.Parse(dbUrl)
		if err != nil {
			log.Fatalf(`could not parse %s`, dbUrl)
		}

		host, _, err := net.SplitHostPort(u.Host)
		if err != nil {
			log.Fatalf(`could not split host port %s`, u.Host)
		}
		u.Host = host + ":5433"
		dbUrl = u.String()
	}

	// db
	db, err := pgxpool.Connect(context.Background(), dbUrl)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// routes
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var col int
		db.QueryRow(r.Context(), "SELECT 1").Scan(&col)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, "{\"status\":\"ok\"}")
	})

	// listen
	log.Println("Listening at http://localhost:" + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
