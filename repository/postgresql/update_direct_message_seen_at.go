package postgresql

import (
	"database/sql"
	"senggol/model"
)

type UpdateDirectMessageSeenAtRepository struct {
	db *sql.DB
}

func (repo UpdateDirectMessageSeenAtRepository) UpdateDirectMessageSeenAt(id int, seenAt time.Time) err
	return nil
}