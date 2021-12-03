package postgresql

import (
	"database/sql"
	"senggol/model"
)

type GetPeerDirectMessagesRepository struct {
	db *sql.DB
}

func (repo GetPeerDirectMessagesRepository) GetPeerDirectMessages(id, prev, limit int) ([]model.DirectMessageWithContent, err) {
	return []model.DirectMessageWithContent{}, nil
}