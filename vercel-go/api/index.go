package api

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	dbUrl, ok := os.LookupEnv("DATABASE_URL")
	if !ok {
		log.Fatal("DATABASE_URL not set")
		fmt.Fprintf(w, "{\"status\":\"internal server error\"}")
		return
	}

	db, err := sql.Open("postgres", dbUrl)
	if err != nil {
		log.Fatal(err)
		fmt.Fprintf(w, "{\"status\":\"internal server error\"}")
		return
	}
	defer db.Close()

	rows, err := db.Query("SELECT 1")
	if err != nil {
		log.Fatal(err)
		fmt.Fprintf(w, "{\"status\":\"internal server error\"}")
		return
	}
	defer rows.Close()

	fmt.Fprintf(w, "{\"status\":\"ok\"}")
}
