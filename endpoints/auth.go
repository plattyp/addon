package endpoints

import (
	"github.com/gin-gonic/gin"
	"github.com/plattyp/addon/resources"
)

// Signup allows you to create a new User
func Signup(c *gin.Context) {
	var json resources.Auth
	err := c.BindJSON(&json)
	if err == nil {
		Success("Authenticated successfully", c)
	}
}
