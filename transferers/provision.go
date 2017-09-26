package transferers

// Provision is the binding for the Heroku endpoint
type Provision struct {
	HerokuID    string `form:"heroku_id" json:"heroku_id" binding:"required"`
	Plan        string `form:"plan" json:"plan" binding:"required"`
	Region      string `form:"region" json:"region"`
	CallbackURL string `form:"callback_url" json:"callback_url"`
}

// GetValues returns back a map of values about the Provision transferer
func (p Provision) GetValues() map[string]interface{} {
	return map[string]interface{}{
		"heroku_id":    p.HerokuID,
		"plan":         p.Plan,
		"region":       p.Region,
		"callback_url": p.CallbackURL,
	}
}
