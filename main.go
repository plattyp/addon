package main

import (
	"github.com/plattyp/addon/endpoints"
	"github.com/plattyp/addon/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Use(middleware.Errors())
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.GET("/", endpoints.Index)

	auth := router.Group("/auth")
	{
		auth.POST("/signup", endpoints.Signup)
	}

	router.POST("/heroku/resources", endpoints.HerokuProvision)

	// Generic 404
	router.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"status": false, "message": "Resource not found"})
	})

	router.Run(":5000")
}
