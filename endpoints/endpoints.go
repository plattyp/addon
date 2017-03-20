package endpoints

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Success returns generic success message
func Success(message string, c *gin.Context) {
	c.JSON(
		http.StatusCreated,
		gin.H{
			"status":  true,
			"message": message,
		},
	)
}
