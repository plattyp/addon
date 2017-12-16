package accessor

import (
	"time"

	"github.com/lib/pq"
)

func getDeletedAtField() map[string]interface{} {
	return map[string]interface{}{
		"deleted_at": pq.NullTime{Time: time.Now().UTC(), Valid: true},
	}
}
