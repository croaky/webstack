package api

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Server struct {
	db *pgxpool.Pool
}

func NewServer(db *pgxpool.Pool) *Server {
	return &Server{
		db: db,
	}
}

func (s *Server) Serve(port string) {
	r := chi.NewRouter()

	// middleware
	r.Use(middleware.RequestID)
	r.Use(Logger())
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(20 * time.Second))
	r.Use(Headers)

	// OPTIONS catch-all for pre-flight requests
	r.Options("/*", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(204)
		w.Write([]byte(""))
	})

	// routes
	r.Get("/", s.healthHandler)

	fmt.Println("Listening on http://localhost:" + port)
	http.ListenAndServe(":"+port, r)
}

func Headers(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		w.Header().Add("Accept", "application/json")
		w.Header().Add("Access-Control-Allow-Methods", "OPTIONS,POST")
		w.Header().Add("Content-Type", "application/json")
		w.Header().Add("X-Content-Type-Options", "nosniff")
		w.Header().Add("X-Frame-Options", "DENY")
		w.Header().Add("X-XSS-Protection", "1; mode=block")
		next.ServeHTTP(w, r.WithContext(ctx))
	}
	return http.HandlerFunc(fn)
}

func (s *Server) healthHandler(w http.ResponseWriter, r *http.Request) {
	var col int
	s.db.QueryRow(r.Context(), "SELECT 1").Scan(&col)
	fmt.Fprintf(w, "{\"status\":\"ok\"}")
}
