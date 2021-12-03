package postgresql

import (
	"database/sql"
	"senggol/model"
)

type GetPeersCountRepository struct {
	db *sql.DB
}

func (repo GetPeersCountRepository) GetPeersCount(id int) (int, err) {
	var count int
	return count, nil
}