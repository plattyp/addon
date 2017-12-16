package endpoints

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/plattyp/addon/accessor"
	"github.com/plattyp/addon/transferers"
)

// HerokuProvision is for provisioning the Heroku Addon
func (e *Endpointer) HerokuProvision(c *gin.Context) {
	var json transferers.Provision

	err := c.ShouldBindWith(&json, binding.JSON)
	if err == nil {
		successHerokuProvision(&json, e, c)
	} else {
		processError(err, c)
	}
}

func successHerokuProvision(json *transferers.Provision, e *Endpointer, c *gin.Context) {
	// Lookup Plan via Code
	p := accessor.PlanDataAccessor{Databaser: e.databaser}
	plan, lErr := p.GetPlanByCode(json.Plan)
	if lErr != nil {
		Error(lErr.Error(), c)
		return
	}

	// Create User And Associate With Plan
	u := accessor.UserDataAccessor{Databaser: e.databaser}
	user, uErr := u.CreateUser(plan.ID, json.Region, json.HerokuID)
	if uErr != nil {
		Error(uErr.Error(), c)
		return
	}

	// Create App for User
	a := accessor.AppDataAccessor{Databaser: e.databaser}
	_, aErr := a.CreateApp(user.ID)
	if aErr != nil {
		Error(aErr.Error(), c)
		return
	}

	// Create the response
	response := transferers.NewHerokuResource(user.ID, "Successfully provisioned addon")
	c.JSON(
		http.StatusCreated,
		response,
	)
}
