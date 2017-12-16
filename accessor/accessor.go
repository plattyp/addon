package accessor

import (
	"time"

	"github.com/lib/pq"
	upper "upper.io/db.v3"
)

func getDeletedAtField() map[string]interface{} {
	return map[string]interface{}{
		"deleted_at": pq.NullTime{Time: time.Now().UTC(), Valid: true},
	}
}

func deleteWithCondition(collection upper.Collection, condition upper.Cond) error {
	deletedAt := getDeletedAtField()
	err := collection.Find(condition).Update(&deletedAt)
	if err != nil {
		return err
	}
	return nil
}
