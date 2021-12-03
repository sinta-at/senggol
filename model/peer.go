package model

import (
	"time"
)

type Peer struct {
	ID                  string    `json:"id"`
	Username            string    `json:"username"`
	LatestMessageAt     time.Time `json:"latest_message_at"`
	UnseenMessagesCount int       `json:"unseen_messages_count"`
}