package resources

// App is the struct around the Apps table
type App struct {
	Resource
	Name   string `db:"name,omitempty"`
	UserID int64  `db:"user_id"`
}

// GetValues returns back a map of values about the App resource
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

// Table returns the table associated with the resource
func (a App) Table() string {
	return "apps"
}
