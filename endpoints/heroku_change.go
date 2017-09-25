package endpoints

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/plattyp/addon/accessor"
	"github.com/plattyp/addon/transferers"
)

// HerokuChange is for updating existing Heroku addons
func (e *Endpointer) HerokuChange(c *gin.Context) {
	userID := c.Param("id")

	var json transferers.Provision

	err := c.ShouldBindWith(&json, binding.JSON)
	if err == nil {
		// Convert User ID To Usable Int64
		id, iErr := strconv.ParseInt(userID, 10, 64)
		if err != nil {
			Error(iErr.Error(), c)
			return
		}

		fmt.Println(json)

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

		Success("Successfully changed plans", c)
	} else {
		validationErr := HandleError(err)
		ValidationError(validationErr, c)
	}
}
