package postgresql

import (
	"database/sql"
	"time"
)

type UpdateDirectMessageSeenAtRepository struct {
	db *sql.DB
}

func (repo UpdateDirectMessageSeenAtRepository) UpdateDirectMessageSeenAt(messageID int, seenAt time.Time) error {
	return nil
}