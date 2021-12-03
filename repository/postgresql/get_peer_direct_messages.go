package postgresql

import (
	"fmt"
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
	AND direct_messages.peer_id = $2 %s
	ORDER by messages.created_at DESC
	LIMIT $3;`

	cursorFilter := ""
	if prev > 0 {
		cursorFilter = fmt.Sprintf("AND direct_messages.id < %v", prev)
	}

	peerDirectMessages := []model.PeerDirectMessage{}

	rows, err := repo.db.Query(fmt.Sprintf(dql, cursorFilter), userID, peerID, limit)
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