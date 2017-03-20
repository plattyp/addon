package endpoints

import (
	"github.com/gin-gonic/gin"
	"github.com/plattyp/addon/resources"
)

// HerokuProvision is for provisioning the Heroku Addon
func HerokuProvision(c *gin.Context) {
	var json resources.Provision
	err := c.BindJSON(&json)
	if err == nil {
		Success("Addon provisioned successfully.", c)
	}
}
