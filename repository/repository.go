package repository

import (
	"senggol/model"

	"time"
)

type CreateUser interface {
	CreateUser(user model.User) error
}

type GetUserByUsername interface {
	GetUserByUsername(username string) (model.User, error)
}

type GetPeers interface {
	GetPeers(userID, limit, offset int) ([]model.Peer, error)
}

type GetPeersCount interface {
	GetPeersCount(userID int) (int, error)
}

type GetPeerDirectMessages interface {
	GetPeerDirectMessages(userID, peerID, prev, limit int) ([]model.PeerDirectMessage, error)
}

type CreateDirectMessage interface {
	CreateDirectMessage(message model.Message, inboundDirectMessage, outboundDirectMessage model.DirectMessage) error
}

type UpdateDirectMessageSeenAt interface {
	UpdateDirectMessageSeenAt(messageID int, seenAt time.Time) error
}