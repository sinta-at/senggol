package postgresql

import (
	"database/sql"
	"senggol/model"
)

type GetPeersRepository struct {
	db *sql.DB
}

func (repo GetPeersRepository) GetPeers(id, limit, offset int) ([]model.Peer, error) {
	return []model.Peer{}, nil
}