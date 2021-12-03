package postgresql

import (
	"database/sql"
)

type GetPeersCountRepository struct {
	db *sql.DB
}

func (repo GetPeersCountRepository) GetPeersCount(userID int) (int, error) {
	dql := `SELECT COUNT(*) FROM users where id != $1`
	count := 0
	if err := repo.db.QueryRow(dql, userID).Scan(&count); err != nil {
		return 0, err
	}
	return count, nil
}