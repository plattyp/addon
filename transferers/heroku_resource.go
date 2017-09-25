package transferers

// HerokuResource is the binding for responding via the Heroku endpoint
type HerokuResource struct {
	Id      int64  `json:"id"`
	Message string `json:"message"`
}

func NewHerokuResource(id int64, message string) *HerokuResource {
	return &HerokuResource{
		Id:      id,
		Message: message,
	}
}
