package transferers

// HerokuResource is the binding for responding via the Heroku endpoint
type HerokuResource struct {
	ID      int64  `json:"id"`
	Message string `json:"message"`
}

// NewHerokuResource returns back a HerokuResource
func NewHerokuResource(id int64, message string) *HerokuResource {
	return &HerokuResource{
		ID:      id,
		Message: message,
	}
}
