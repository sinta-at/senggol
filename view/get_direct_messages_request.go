package view

type GetDirectMessagesRequest struct {
	UserID int
	PeerID int
	Prev   int
	Limit  int
}