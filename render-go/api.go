package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4/pgxpool"
)

func env(key string, fallback string) string {
	val := os.Getenv(key)
	if val == "" {
		val = fallback
	}
	return val
}

type Server struct {
	db *pgxpool.Pool
	rt *gin.Engine
}

func NewServer(ctx context.Context, db *pgxpool.Pool) *Server {
	gin.SetMode(gin.ReleaseMode)
	rt := gin.New()

	rt.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("%s %d %s %s\n",
			fmt.Sprintf("% 4dms", param.Latency.Milliseconds()),
			param.StatusCode,
			param.Method,
			param.Path,
			// TODO: req_id, git_sha
		)
	}))
	rt.Use(gin.Recovery())

	rt.GET("/", func(c *gin.Context) {
		var col int
		db.QueryRow(ctx, "SELECT 1").Scan(&col)
		c.JSON(200, gin.H{"status": "ok"})
	})

	return &Server{
		db: db,
		rt: rt,
	}
}

func main() {
	ctx := context.Background()

	// env
	port := env("PORT", "8080")
	dbUrl := env("DATABASE_URL", "postgres:///webstack_dev")

	// db
	db, err := pgxpool.Connect(ctx, dbUrl)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// server
	s := NewServer(ctx, db)

	log.Println("Listening at http://localhost:" + port)
	s.rt.Run(":" + port)
}
