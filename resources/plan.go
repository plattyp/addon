package resources

// Plan is the struct around the Plans table
type Plan struct {
	Resource
	Code        string `db:"code"`
	Name        string `db:"name"`
	Description string `db:"description"`
	Ordinal     int    `db:"ordinal"`
}

// GetValues returns back a map of values about the Plan resource
func (plan Plan) GetValues() map[string]interface{} {
	return map[string]interface{}{
		"id":          plan.ID,
		"code":        plan.Code,
		"name":        plan.Name,
		"description": plan.Description,
		"ordinal":     plan.Ordinal,
		"created_at":  plan.CreatedAt,
		"updated_at":  plan.UpdatedAt,
		"deleted_at":  plan.DeletedAt,
	}
}

// Table returns the table associated with the resource
func (plan Plan) Table() string {
	return "plans"
}
