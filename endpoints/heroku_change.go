package endpoints

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/plattyp/addon/accessor"
	"github.com/plattyp/addon/transferers"
)

// HerokuChange is for updating existing Heroku addons
func (e *Endpointer) HerokuChange(c *gin.Context) {
	var json transferers.Provision

	err := BindJSONToTransferer(&json, c)
	if err == nil {
		successHerokuChange(&json, e, c)
	} else {
		processError(err, c)
	}
}

func successHerokuChange(json *transferers.Provision, e *Endpointer, c *gin.Context) {
	userID := c.Param("id")

	// Convert User ID To Usable Int64
	id, iErr := strconv.ParseInt(userID, 10, 64)
	if iErr != nil {
		Error(iErr.Error(), c)
		return
	}

	// Lookup Plan via Code
	p := accessor.PlanDataAccessor{Databaser: e.databaser}
	plan, lErr := p.GetPlanByCode(json.Plan)
	if lErr != nil {
		Error(lErr.Error(), c)
		return
	}

	// Update Plan For User
	u := accessor.UserDataAccessor{Databaser: e.databaser}
	pErr := u.UpdatePlan(id, plan.ID)
	if pErr != nil {
		Error(pErr.Error(), c)
		return
	}

	SuccessOK("Successfully changed plans", c)
}
