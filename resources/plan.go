package resources

// Plan is the struct around the Users table
type Plan struct {
	Resource
	Code        string `db:"code"`
	Name        string `db:"name"`
	Description string `db:"description"`
	Ordinal     int    `db:"ordinal"`
}

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

func (plan Plan) Table() string {
	return "plans"
}
