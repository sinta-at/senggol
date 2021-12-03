package postgresql

import (
	"database/sql"
	"senggol/model"
)

type GetUserByUsernameRepository struct {
	db *sql.DB
}

func (repo GetUserByUsernameRepository) GetUserByUsername(username string) (model.User, error) {
	return model.User{}, nil
}