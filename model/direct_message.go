package model

import (
	"time"
)

type DirectMessage struct {
	ID        int
	UserID    int
	PeerID    int
	MessageID int
	Direction string
	SeenAt    *time.Time
}

type PeerDirectMessage struct {
	ID             int
	PeerID         int
	MessageID      int
	MessageContent string
	Direction      string
	SeenAt         *time.Time
	CreatedAt      time.Time
}