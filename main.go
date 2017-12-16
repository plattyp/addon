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
	// Create a DB Connection
	dbConn, err := db.NewDatabaser(os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal("NewDatabaser: ", err)
	}

	// Create an Endpointer
	e := endpoints.NewEndpointer(dbConn)

	router := setupRouter(e)
	router.Run(":5000")
}

func loadEnvironment() {
	addonEnv := os.Getenv("ADDON_ENVIRONMENT")

	// Load from .env if development or travis
	if addonEnv == "" || addonEnv == "development" || addonEnv == "travis" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}
}

func setupRouter(e *endpoints.Endpointer) *gin.Engine {
	router := gin.Default()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.LoadHTMLGlob("templates/*")

	router.GET("/", endpoints.Index)

	herokuAuthorized := router.Group("/heroku", gin.BasicAuth(gin.Accounts{
		os.Getenv("HEROKU_USERNAME"): os.Getenv("HEROKU_PASSWORD"),
	}))

	herokuAuthorized.POST("/resources", e.HerokuProvision)
	herokuAuthorized.PUT("/resources/:id", e.HerokuChange)
	herokuAuthorized.DELETE("/resources/:id", e.HerokuDelete)

	herokoUnauthorized := router.Group("/heroku")
	herokoUnauthorized.POST("/sso", e.HerokuSSO)

	// Generic 404
	router.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"status": false, "message": "Resource not found"})
	})

	return router
}
