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
	ID             int        `json:"id"`
	PeerID         int        `json:"peer_id"`
	MessageID      int        `json:"message_id"`
	MessageContent string     `json:"message_content"`
	Direction      string     `json:"direction"`
	SeenAt         *time.Time `json:"seen_at,omitempty"`
	CreatedAt      time.Time  `json:"created_at"`
}