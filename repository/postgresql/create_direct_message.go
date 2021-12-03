package postgresql

import (
	"database/sql"
	"senggol/model"
)

type CreateDirectMessageRepository struct {
	db *sql.DB
}

func (repo CreateDirectMessageRepository) CreateDirectMessage(message model.Message, directMessage model.DirectMessage) err {
	return nil
}