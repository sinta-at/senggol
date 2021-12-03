package repository

import (
	"senggol/model"

	"time"
)

type CreateUser interface {
	CreateUser(model.User) err
}

type GetUserByID interface {
	GetUserByID(id int) (model.User, err)
}

type GetPeers interface {
	GetPeers(id, limit, offset int) ([]model.Peer, err)
}

type GetPeersCount interface {
	GetPeersCount(id int) (int, err)
}

type GetPeerDirectMessages interface {
	GetPeerDirectMessages(id, prev, limit int) ([]model.DirectMessageWithContent, err)
}

type CreateDirectMessage interface {
	CreateDirectMessage(message model.Message, directMessage model.DirectMessage) err
}

type UpdateDirectMessageSeenAt interface {
	UpdateDirectMessageSeenAt(id int, seenAt time.Time) err
}