package main

import (
	"database/sql"
	_ "embed"
	"flag"
	"fmt"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// flag
	dsn := flag.String("dsn", "", "datasource name")
	flag.Parse()
	if *dsn == "" {
		log.Fatalf("dsn required")
	}

	// db
	db, err := sql.Open("sqlite3", *dsn)
	if err != nil {
		log.Fatalf("open db: %v", err)
	}
	defer db.Close()

	// route
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		db.Query("SELECT 1")
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, "{\"status\":\"ok\"}")
	})

	log.Println("http server listening on :8080")
	http.ListenAndServe(":8080", nil)
}
