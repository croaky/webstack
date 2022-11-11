package main

import (
	"context"
	"embed"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
)

var (
	//go:embed html/*.html
	views embed.FS

	html = template.Must(template.ParseFS(views, "html/*.html"))
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
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// routes
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		err := html.ExecuteTemplate(w, "index.html", nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	http.HandleFunc("/clicked", func(w http.ResponseWriter, r *http.Request) {
		var version string
		db.QueryRow(r.Context(), "SELECT version()").Scan(&version)
		resp := struct {
			Version string
		}{
			Version: version,
		}
		err := html.ExecuteTemplate(w, "version.html", resp)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	// listen
	log.Println("Listening at http://localhost:" + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
