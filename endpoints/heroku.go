package endpoints

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/plattyp/addon/accessor"
)

// Provision is the binding for the Heroku endpoint
type Provision struct {
	HerokuID    string `form:"heroku_id" json:"heroku_id" binding:"required"`
	Plan        string `form:"plan" json:"plan" binding:"required"`
	Region      string `form:"region" json:"region" binding:"required"`
	CallbackURL string `form:"callback_url" json:"callback_url" binding:"required"`
}

// GetValues returns back a map of all values of the given resource
func (p Provision) GetValues() map[string]interface{} {
	return map[string]interface{}{
		"heroku_id":    p.HerokuID,
		"plan":         p.Plan,
		"region":       p.Region,
		"callback_url": p.CallbackURL,
	}
}

// HerokuProvision is for provisioning the Heroku Addon
func (e *Endpointer) HerokuProvision(c *gin.Context) {
	var json Provision

	err := c.ShouldBindWith(&json, binding.JSON)
	if err == nil {
		fmt.Println(json.GetValues())

		// Lookup Plan via Code
		p := accessor.PlanDataAccessor{Databaser: e.databaser}
		plan, lErr := p.GetPlanByCode(json.Plan)
		if err != nil {
			Error(lErr.Error(), c)
			return
		}

		// Create User And Associate With Plan

		// Create App for User

		// Return User ID

		fmt.Println(plan)
		// Actually perform work and return resource
		Success("Addon provisioned successfully.", c)
	} else {
		validationErr := HandleError(err)
		ValidationError(validationErr, c)
	}
}
