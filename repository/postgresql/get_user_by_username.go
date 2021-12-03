package postgresql

import (
	"database/sql"
	"senggol/model"
)

type GetUserByUsernameRepository struct {
	db *sql.DB
}

func (repo GetUserByUsernameRepository) GetUserByUsername(username string) (model.User, error) {
	var user model.User

	dql := `SELECT
		id, username, password, created_at
	FROM users
	WHERE username = $1`

	err := repo.db.QueryRow(dql, username).Scan(&user.ID, &user.Username, &user.Password, &user.CreatedAt)
	if err != nil {
		return model.User{}, err
	}

	return user, nil
}