package accessor

import (
	"database/sql"
	"errors"
	"time"

	"github.com/lib/pq"
	"github.com/plattyp/addon/db"
	"github.com/plattyp/addon/resources"
	upper "upper.io/db.v3"
)

// UserAccessor is used to fetch users
type UserAccessor interface {
	CreateUser(planID int64, region string, herokuID string) (*resources.User, error)
	FetchUser(id int64) (*resources.User, error)
	UpdatePlan(id int64, planID int64) error
	DeleteUser(id int64) error
}

const userTableName = "users"

// UserDataAccessor is used to interacts with users
type UserDataAccessor struct {
	Databaser *db.Databaser
}

// CreateUser will create a user based on the params
func (u UserDataAccessor) CreateUser(planID int64, region string, herokuID string) (*resources.User, error) {
	// Create User Associated With Plan
	user := resources.User{
		PlanID:     planID,
		Region:     sql.NullString{String: region, Valid: true},
		HerokuUUID: sql.NullString{String: herokuID, Valid: true},
	}

	result, err := u.usersTable().Insert(&user)
	if err != nil {
		return nil, err
	}

	if id, ok := result.(int64); ok {
		user, err := u.FetchUser(id)
		if err != nil {
			return nil, err
		}
		return user, nil
	}

	return nil, errors.New("Failed to create user")
}

// UpdatePlan updates the plan associated with the user
func (u UserDataAccessor) UpdatePlan(id int64, planID int64) error {
	var user = resources.User{
		PlanID: planID,
	}

	err := u.usersTable().Find(upper.Cond{"id": id}).Update(&user)
	if err == nil {
		return nil
	}

	return err
}

// FetchUser retrieves a user by their ID
func (u UserDataAccessor) FetchUser(id int64) (*resources.User, error) {
	var user resources.User

	err := u.usersTable().Find(upper.Cond{"id": id}).One(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// DeleteUser deletes the associated user
func (u UserDataAccessor) DeleteUser(id int64) error {
	deletedAt := map[string]interface{}{
		"deleted_at": pq.NullTime{Time: time.Now().UTC(), Valid: true},
	}
	err := u.usersTable().Find(upper.Cond{"id": id}).Update(&deletedAt)
	if err == nil {
		return nil
	}

	return err
}

// usersTable returns back a collection
func (u UserDataAccessor) usersTable() upper.Collection {
	return u.Databaser.Conn.Collection(userTableName)
}
