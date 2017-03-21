package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/plattyp/addon/endpoints"
	"github.com/plattyp/addon/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	router := gin.Default()
	router.Use(middleware.ValidateErrors())
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.GET("/", endpoints.Index)

	auth := router.Group("/auth")
	{
		auth.POST("/signup", endpoints.Signup)
	}

	herokuAuthorized := router.Group("/heroku", gin.BasicAuth(gin.Accounts{
		os.Getenv("HEROKU_USERNAME"): os.Getenv("HEROKU_PASSWORD"),
	}))

	herokuAuthorized.POST("/resources", endpoints.HerokuProvision)

	// Generic 404
	router.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"status": false, "message": "Resource not found"})
	})

	router.Run(":5000")
}
