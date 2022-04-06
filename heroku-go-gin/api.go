package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/gin-gonic/gin"

	_ "github.com/lib/pq"
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

	// db conn
	db, err := sql.Open("postgres", dbUrl)
	if err != nil {
		log.Fatal(err)
	}

	// router
	router := gin.New()
	router.Use(gin.Logger())

	router.GET("/", func(ctx *gin.Context) {
		db.Query("SELECT 1")

		ctx.JSON(200, gin.H{
			"status": "ok",
		})
	})

	// serve
	router.Run(":" + port)
}
