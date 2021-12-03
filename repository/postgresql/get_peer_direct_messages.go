package postgresql

import (
	"database/sql"
	"senggol/model"
)

type GetPeerDirectMessagesRepository struct {
	db *sql.DB
}

func (repo GetPeerDirectMessagesRepository) GetPeerDirectMessages(userID, peerID, prev, limit int) ([]model.PeerDirectMessage, error) {
	return []model.PeerDirectMessage{}, nil
}