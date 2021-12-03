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
	SeenAt    time.Time
}

type DirectMessageWithContent struct {
	ID             int
	UserID         int
	PeerID         int
	MessageID      int
	MessageContent string
	Direction      string
	SeenAt         time.Time
}