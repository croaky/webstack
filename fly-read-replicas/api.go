package main

import (
	"database/sql"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/url"
	"os"

	_ "github.com/lib/pq"
)

func main() {
	// env
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	dbUrl := os.Getenv("DATABASE_URL")
	if dbUrl == "" {
		dbUrl = "postgres://postgres:postgres@localhost:5432/webstack_dev"
	}
	primary := os.Getenv("PRIMARY_REGION")
	current := os.Getenv("FLY_REGION")
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
	log.Println("Connecting to database at " + dbUrl)
	db, err := sql.Open("postgres", dbUrl)
	if err != nil {
		log.Fatal(err)
	}

	// routes
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		db.Query("SELECT 1")
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, "{\"status\":\"ok\"}")
	})

	// listen
	log.Println("Listening at http://localhost:" + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
