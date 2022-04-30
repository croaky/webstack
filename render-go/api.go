package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
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
	ctx := context.Background()

	// db
	dbpool, err := pgxpool.Connect(ctx, os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	defer dbpool.Close()

	// init router
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()

	// middleware
	r.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("method=%s path=%s status=%d duration=%s error=%s\n",
			param.Method,
			param.Path,
			param.StatusCode,
			fmt.Sprintf("%dms", param.Latency.Milliseconds()),
			param.ErrorMessage,
			// TODO: req_id=, req_ip=, release=, git_head=, user_id=, err_id=,
			// rate_limit_enforced=false, rate_limit=100=, rate_limit_remaining=99
			// https://brandur.org/canonical-log-lines
		)
	}))
	r.Use(gin.Recovery())

	r.GET("/", func(c *gin.Context) {
		dbpool.QueryRow(ctx, "SELECT 1")
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})

	// serve
	log.Println("Listening at http://localhost:" + port)
	r.Run(":" + port)
}
