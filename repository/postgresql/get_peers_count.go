package postgresql

import (
	"database/sql"
)

type GetPeersCountRepository struct {
	db *sql.DB
}

func (repo GetPeersCountRepository) GetPeersCount(id int) (int, error) {
	var count int
	return count, nil
}