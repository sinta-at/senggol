package postgresql

import (
	"database/sql"
	"senggol/model"
)

type GetUserByIDRepository struct {
	db *sql.DB
}

func (repo GetUserByIDRepository) GetUserByID(id int) (model.User, err) {
	return model.User{}, nil
}