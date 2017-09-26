package endpoints

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/plattyp/addon/accessor"
)

// HerokuDelete is for updating existing Heroku addons
func (e *Endpointer) HerokuDelete(c *gin.Context) {
	userID := c.Param("id")

	// Convert User ID To Usable Int64
	id, iErr := strconv.ParseInt(userID, 10, 64)
	if iErr != nil {
		Error(iErr.Error(), c)
		return
	}

	// Fetch User to ensure it exists
	u := accessor.UserDataAccessor{Databaser: e.databaser}
	user, uErr := u.FetchUser(id)
	if uErr != nil {
		Error(uErr.Error(), c)
		return
	}

	// Delete Apps Associated With User
	a := accessor.AppDataAccessor{Databaser: e.databaser}
	aErr := a.DeleteAppsByUser(user.ID)
	if aErr != nil {
		Error(aErr.Error(), c)
		return
	}

	// Delete User
	dErr := u.DeleteUser(user.ID)
	if dErr != nil {
		Error(dErr.Error(), c)
		return
	}

	SuccessOK("Successfully deleted user.", c)
}
