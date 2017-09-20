package resources

import (
	"time"

	"github.com/lib/pq"
)

// Resourcer is a generic resource interface
type Resourcer interface {
	GetValues() map[string]interface{}
	TableName() string
}

// Resource contains all basic functionality that all database models share
type Resource struct {
	Resourcer
	ID        uint        `db:"id,omitempty"`
	CreatedAt time.Time   `db:"created_at"`
	UpdatedAt time.Time   `db:"updated_at"`
	DeletedAt pq.NullTime `db:"deleted_at"`
}

// GetValues returns back a map of interfaces of the value of the resource
func (resource Resource) GetValues() map[string]interface{} {
	return map[string]interface{}{
		"id":         resource.ID,
		"created_at": resource.CreatedAt,
		"updated_at": resource.UpdatedAt,
		"deleted_at": resource.DeletedAt,
	}
}

// TableName returns back the name of the table within the database
func (resource Resource) TableName() string {
	return ""
}
