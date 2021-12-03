package repository

import (
	"senggol/model"

	"time"
)

type CreateUser interface {
	CreateUser(model.User) error
}

type GetUserByUsername interface {
	GetUserByUsername(username string) (model.User, error)
}

type GetPeers interface {
	GetPeers(id, limit, offset int) ([]model.Peer, error)
}

type GetPeersCount interface {
	GetPeersCount(id int) (int, error)
}

type GetPeerDirectMessages interface {
	GetPeerDirectMessages(id, prev, limit int) ([]model.DirectMessageWithContent, error)
}

type CreateDirectMessage interface {
	CreateDirectMessage(message model.Message, directMessage model.DirectMessage) error
}

type UpdateDirectMessageSeenAt interface {
	UpdateDirectMessageSeenAt(id int, seenAt time.Time) error
}