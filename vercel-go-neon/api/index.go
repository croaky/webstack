package api

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/jackc/pgx/v4"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	dsn, ok := os.LookupEnv("DATABASE_URL")
	if !ok {
		log.Fatal("DATABASE_URL not set")
		fmt.Fprintf(w, "{\"status\":\"internal server error\"}")
		return
	}

	ctx := context.Background()
	db, err := pgx.Connect(ctx, dsn)
	defer db.Close(ctx)
	if err != nil {
		log.Fatal(err)
		fmt.Fprintf(w, "{\"status\":\"internal server error\"}")
		return
	}

	var col int
	err = db.QueryRow(ctx, "SELECT 1").Scan(&col)
	if err != nil {
		log.Fatal(err)
		fmt.Fprintf(w, "{\"status\":\"internal server error\"}")
		return
	}

	fmt.Fprintf(w, "{\"status\":\"ok\"}")
}
