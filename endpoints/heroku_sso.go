package endpoints

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/plattyp/addon/accessor"
	"github.com/plattyp/addon/services"
	"github.com/plattyp/addon/transferers"
)

// HerokuSSO is for authenticating from Heroku
func (e *Endpointer) HerokuSSO(c *gin.Context) {
	var json transferers.SSO

	err := c.ShouldBindWith(&json, binding.FormPost)
	if err == nil {
		u := accessor.UserDataAccessor{Databaser: e.databaser}
		user, statusCode, aErr := services.ValidateSSOToken(u, json.ID, json.Token, json.Timestamp)
		if aErr != nil {
			c.JSON(
				statusCode,
				gin.H{
					"status":  false,
					"message": aErr.Error(),
				},
			)
			return
		}

		// Add User To Context
		if user != nil {
			fmt.Printf("Authenticated User: %s", user.Email.String)
		}

		// Save Cookie
		c.SetCookie("heroku-nav-data", json.NavData, 3600, "/", "", false, false)

		// Render HTML With Heroku SSO navigation header
		c.HTML(
			statusCode,
			"dashboard.tmpl",
			gin.H{
				"title": "dashboard",
				"app":   json.App,
				"addon": os.Getenv("HEROKU_ADDON_NAME"),
			},
		)
	} else {
		validationErr := HandleError(err)
		ValidationError(validationErr, c)
	}
}
