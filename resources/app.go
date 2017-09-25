package resources

// User is the struct around the Users table
type App struct {
	Resource
	Name   string `db:"name,omitempty"`
	UserID int64  `db:"user_id"`
}

func (a App) GetValues() map[string]interface{} {
	return map[string]interface{}{
		"id":         a.ID,
		"name":       a.Name,
		"user_id":    a.UserID,
		"created_at": a.CreatedAt,
		"updated_at": a.UpdatedAt,
		"deleted_at": a.DeletedAt,
	}
}

func (a App) Table() string {
	return "apps"
}
