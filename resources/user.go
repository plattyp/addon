package resources

// User is the struct around the Users table
type User struct {
	Resource
	Email  string `db:"email"`
	PlanID string `db:"plan_id"`
}

func (user User) GetValues() map[string]interface{} {
	return map[string]interface{}{
		"id":         user.ID,
		"email":      user.Email,
		"plan_id":    user.PlanID,
		"created_at": user.CreatedAt,
		"updated_at": user.UpdatedAt,
		"deleted_at": user.DeletedAt,
	}
}

func (user User) Table() string {
	return "users"
}
