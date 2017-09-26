package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/plattyp/addon/db"
	"github.com/plattyp/addon/endpoints"

	"github.com/gin-gonic/gin"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	router := gin.Default()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.GET("/", endpoints.Index)

	// Create a DB Connection
	dbConn, err := db.NewDatabaser(os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal("NewDatabaser: ", err)
	}

	// Create an Endpointer
	e := endpoints.NewEndpointer(dbConn)

	herokuAuthorized := router.Group("/heroku", gin.BasicAuth(gin.Accounts{
		os.Getenv("HEROKU_USERNAME"): os.Getenv("HEROKU_PASSWORD"),
	}))

	herokuAuthorized.POST("/resources", e.HerokuProvision)
	herokuAuthorized.PUT("/resources/:id", e.HerokuChange)
	herokuAuthorized.DELETE("/resources/:id", e.HerokuDelete)

	// Generic 404
	router.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"status": false, "message": "Resource not found"})
	})

	router.Run(":5000")
}
