package view

import (
	"senggol/model"
)

type GetDirectMessagesResponse struct {
	Messages       []model.PeerDirectMessage `json:"messages"`
	CursorPosition int                       `json:"cursor_position"`
}