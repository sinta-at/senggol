package postgresql

import (
	"database/sql"
	"senggol/model"
)

type GetPeerDirectMessagesRepository struct {
	db *sql.DB
}

func (repo GetPeerDirectMessagesRepository) GetPeerDirectMessages(userID, peerID, prev, limit int) ([]model.PeerDirectMessage, error) {
	dql := `SELECT
		direct_messages.id,
		direct_messages.peer_id,
		messages.id,
		messages.content,
		direct_messages.direction,
		direct_messages.seen_at,
		messages.created_at
	FROM direct_messages
	JOIN messages ON messages.id = direct_messages.message_id
	WHERE direct_messages.user_id = $1
	AND direct_messages.peer_id = $2
	AND direct_messages.id < $3
	ORDER by messages.created_at DESC
	LIMIT $4;`

	peerDirectMessages := []model.PeerDirectMessage{}

	rows, err := repo.db.Query(dql, userID, peerID, prev, limit)
	if err != nil {
		return peerDirectMessages, err
	}

	for rows.Next() {
		peerDirectMessage := model.PeerDirectMessage{}
		err := rows.Scan(
			&peerDirectMessage.ID,
			&peerDirectMessage.PeerID,
			&peerDirectMessage.MessageID,
			&peerDirectMessage.MessageContent,
			&peerDirectMessage.Direction,
			&peerDirectMessage.SeenAt,
			&peerDirectMessage.CreatedAt,
		)
		if err != nil {
			return []model.PeerDirectMessage{}, err
		}
		peerDirectMessages = append(peerDirectMessages, peerDirectMessage)
	}

	return peerDirectMessages, nil
}