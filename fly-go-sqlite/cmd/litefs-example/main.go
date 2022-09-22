package main

import (
	"context"
	"database/sql"
	_ "embed"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

// Command line flags.
var (
	dsn  = flag.String("dsn", "", "datasource name")
	addr = flag.String("addr", ":8080", "bind address")
)

var db *sql.DB

func main() {
	log.SetFlags(0)
	rand.Seed(time.Now().UnixNano())

	if err := run(context.Background()); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(ctx context.Context) (err error) {
	// env
	flag.Parse()

	if *dsn == "" {
		return fmt.Errorf("dsn required")
	} else if *addr == "" {
		return fmt.Errorf("bind address required")
	}

	// db
	db, err = sql.Open("sqlite3", *dsn)
	if err != nil {
		return fmt.Errorf("open db: %w", err)
	}
	defer db.Close()

	log.Printf("database opened at %s", *dsn)

	// routes
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		db.Query("SELECT 1")
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, "{\"status\":\"ok\"}")
	})

	log.Printf("http server listening on %s", *addr)
	return http.ListenAndServe(*addr, nil)
}
