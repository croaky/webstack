package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
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

	// db
	db, err := pgxpool.Connect(context.Background(), dbUrl)
	defer db.Close()
	if err != nil {
		log.Fatal(err)
	}

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
