package postgresql

import (
	"database/sql"
	"senggol/model"

	"github.com/pkg/errors"
)

type CreateDirectMessageRepository struct {
	db *sql.DB
}

func (repo CreateDirectMessageRepository) CreateDirectMessage(message model.Message, inboundDirectMessage, outboundDirectMessage model.DirectMessage) error {
	messageDml := `INSERT INTO public.messages
		(id, content, created_at)
		VALUES
		($1, $2, $3);`
	directMessageDml := `INSERT INTO public.direct_messages
		(id, user_id, peer_id, message_id, direction)
		VALUES
		($1, $2, $3, $4, $5),
		($6, $7, $8, $9, $10);`

	dbTx, err := repo.db.Begin()
	if err != nil {
		return err
	}

	_, err = dbTx.Exec(messageDml, message.ID, message.Content, message.CreatedAt)
	if err != nil {
		rollbackErr := dbTx.Rollback()
		if rollbackErr != nil {
			return errors.Wrap(rollbackErr, err.Error())
		}
	}

	_, err = dbTx.Exec(directMessageDml,
		inboundDirectMessage.ID,
		inboundDirectMessage.UserID,
		inboundDirectMessage.PeerID,
		inboundDirectMessage.MessageID,
		inboundDirectMessage.Direction,
		outboundDirectMessage.ID,
		outboundDirectMessage.UserID,
		outboundDirectMessage.PeerID,
		outboundDirectMessage.MessageID,
		outboundDirectMessage.Direction,
	)
	if err != nil {
		rollbackErr := dbTx.Rollback()
		if rollbackErr != nil {
			return errors.Wrap(rollbackErr, err.Error())
		}
	}

	err = dbTx.Commit()
	if err != nil {
		return err
	}

	return nil
}