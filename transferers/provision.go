package transferers

// Provision is the binding for the Heroku endpoint
type Provision struct {
	HerokuID    string `form:"heroku_id" json:"heroku_id" binding:"required"`
	Plan        string `form:"plan" json:"plan" binding:"required"`
	Region      string `form:"region" json:"region" binding:"required"`
	CallbackURL string `form:"callback_url" json:"callback_url" binding:"required"`
}
