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
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	dbUrl := os.Getenv("DATABASE_URL")
	if dbUrl == "" {
		dbUrl = "postgres:///webstack_dev"
	}

	// db
	dbpool, err := pgxpool.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	defer dbpool.Close()

	// routes
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		dbpool.QueryRow(context.Background(), "SELECT 1")
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, "{status:\"ok\"}")
	})

	// serve
	log.Println("Listening at http://localhost:" + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
