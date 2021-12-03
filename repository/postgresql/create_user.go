package postgresql

import (
	"database/sql"
	"senggol/model"
)

type CreateUserRepository struct {
	db *sql.DB
}

func (repo CreateUserRepository) CreateUser(model.User) error {
	return nil
}