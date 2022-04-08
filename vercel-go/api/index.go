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
	dbUrl := os.Getenv("DATABASE_URL")
	db, err := sql.Open("postgres", dbUrl)
	if err != nil {
		log.Fatal(err)
	}
	db.Query("SELECT 1")
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "{status:\"ok\"}")
}
