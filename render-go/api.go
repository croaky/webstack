package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
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
		dbUrl = "postgres:///webstack_dev"
	}

	// db
	db, err := sql.Open("postgres", dbUrl)
	if err != nil {
		log.Fatal(err)
	}

	// routes
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		db.Query("SELECT 1")
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, "{status:\"ok\"}")
	})

	// serve
	log.Println("Listening at http://localhost:" + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
