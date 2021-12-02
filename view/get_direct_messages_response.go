package view

type GetDirectMessagesResponse struct {
	Messages       []DirectMessage `json:"messages"`
	CursorPosition int             `json:"cursor_position"`
}

type DirectMessage struct {
	ID        string `json:"id"`
	Direction string `json:"direction"`
	Content   string `json:"content"`
	SeenAt    string `json:"seen_at"`
	CreatedAt string `json:"created_at"`
}