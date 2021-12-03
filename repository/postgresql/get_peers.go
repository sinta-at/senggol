package postgresql

import (
	"database/sql"
	"senggol/model"
)

type GetPeersRepository struct {
	db *sql.DB
}

func (repo GetPeersRepository) GetPeers(userID, limit, offset int) ([]model.Peer, error) {
	dql := `WITH peer_direct_messages AS (
		SELECT
			direct_messages.id,
			users.id AS peer_id,
			users.username,
			messages.content,
			direct_messages.direction,
			direct_messages.seen_at,
			messages.created_at
		FROM direct_messages
		JOIN messages ON messages.id = direct_messages.id
		JOIN users ON users.id = direct_messages.peer_id
		WHERE direct_messages.user_id = $1
		ORDER BY messages.created_at DESC
	), peer_latest_messages AS (
		SELECT DISTINCT ON (peer_id) peer_id, created_at
		FROM peer_direct_messages
		GROUP BY peer_id, created_at
	), peer_unseen_messages AS (
		SELECT DISTINCT ON (peer_id)
			peer_id,
			count(*) as unread_messages_count
		FROM direct_messages
		WHERE direct_messages.user_id = $1
		AND seen_at IS null
		GROUP BY peer_id
	)
	SELECT
		users.id,
		users.username,
		peer_latest_messages.created_at AS latest_message_at,
		peer_unseen_messages.unread_messages_count
	FROM users
	LEFT JOIN peer_latest_messages ON users.id = peer_latest_messages.peer_id
	LEFT JOIN peer_unseen_messages ON users.id = peer_unseen_messages.peer_id
	WHERE users.id != $1
	ORDER BY peer_latest_messages.created_at DESC
	LIMIT $2 OFFSET $3;`

	peers := []model.Peer{}

	rows, err := repo.db.Query(dql, userID, limit, offset)
	if err != nil {
		return peers, err
	}

	for rows.Next() {
		peer := model.Peer{}
		err := rows.Scan(&peer.ID, &peer.Username, &peer.LatestMessageAt, &peer.UnseenMessagesCount)
		if err != nil {
			return []model.Peer{}, err
		}
		peers = append(peers, peer)
	}
	return peers, nil
}