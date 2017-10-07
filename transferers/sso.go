package transferers

// SSO is the binding for the Heroku SSO endpoint
type SSO struct {
	ID        int64  `form:"id" json:"id" binding:"required"`
	Token     string `form:"token" json:"token" binding:"required"`
	Timestamp int64  `form:"timestamp" json:"timestamp" binding:"required"`
	NavData   string `form:"nav-data" json:"nav-data" binding:"required"`
	Email     string `form:"email" json:"email" binding:"required"`
	App       string `form:"app" json:"app" binding:"required"`
}

// GetValues returns back a map of values about the Provision transferer
func (s SSO) GetValues() map[string]interface{} {
	return map[string]interface{}{
		"id":        s.ID,
		"token":     s.Token,
		"timestamp": s.Timestamp,
		"nav_data":  s.NavData,
		"email":     s.Email,
		"app":       s.App,
	}
}
