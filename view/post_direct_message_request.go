package view

type PostDirectMessageRequest struct {
	UserID  int
	PeerID  int
	Content string `json:"content"`
}