package postgresql

import (
	"database/sql"
	"senggol/model"
)

type CreateUserRepository struct {
	db *sql.DB
}

func (repo CreateUserRepository) CreateUser(user model.User) error {
	dml := `INSERT INTO users
		(id, username, password, created_at)
		VALUES
		($1,$2,$3,$4)`
	if _, err := repo.db.Exec(dml, user.ID, user.Username, user.Password, user.CreatedAt); err != nil {
		return err
	}
	return nil
}