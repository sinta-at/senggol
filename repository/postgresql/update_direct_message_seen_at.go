package postgresql

import (
	"database/sql"
	"time"
)

type UpdateDirectMessageSeenAtRepository struct {
	db *sql.DB
}

func (repo UpdateDirectMessageSeenAtRepository) UpdateDirectMessageSeenAt(messageID int, seenAt time.Time) error {
	dml := `UPDATE direct_messages
		SET seen_at = $1
	WHERE message_id = $2`
	
	if _, err := repo.db.Exec(dml, seenAt, messageID); err != nil {
		return err
	}
	return nil
}