package resources

import "database/sql"

// User is the struct around the Users table
type User struct {
	Resource
	Email      string         `db:"email"`
	PlanID     int64          `db:"plan_id"`
	HerokuUUID sql.NullString `db:"heroku_uuid"`
	Region     sql.NullString `db:"region"`
}

func (user User) GetValues() map[string]interface{} {
	return map[string]interface{}{
		"id":          user.ID,
		"email":       user.Email,
		"plan_id":     user.PlanID,
		"region":      user.Region.String,
		"heroku_uuid": user.HerokuUUID.String,
		"created_at":  user.CreatedAt,
		"updated_at":  user.UpdatedAt,
		"deleted_at":  user.DeletedAt,
	}
}

func (user User) Table() string {
	return "users"
}
