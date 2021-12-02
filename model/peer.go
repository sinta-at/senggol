package model

type Peer struct {
	ID                  string `json:"id"`
	Username            string `json:"username"`
	LatestMessageAt     int    `json:"latest_message_at"`
	UnseenMessagesCount int    `json:"unseen_messages_count"`
}