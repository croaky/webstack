package main

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"

	"server/api"
	"server/env"
)

func main() {
	// env
	dbURL := env.String("DATABASE_URL", "postgres:///webstack_dev")
	port := env.String("PORT", "8080")

	// db
	db, err := pgxpool.Connect(context.Background(), dbURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer db.Close()

	// serve
	server := api.NewServer(db)
	server.Serve(port)
}
