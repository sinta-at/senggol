package service

import (
	"fmt"
	"senggol/enum"
	"senggol/repository"
	"senggol/view"
)

type GetDirectMessagesService struct {
	getPeerDirectMessagesRepository repository.GetPeerDirectMessages
}

func NewGetDirectMessagesService(getPeerDirectMessagesRepository repository.GetPeerDirectMessages) GetDirectMessagesService {
	return GetDirectMessagesService{getPeerDirectMessagesRepository}
}

func (svc GetDirectMessagesService) GetDirectMessages(request view.GetDirectMessagesRequest) (view.GetDirectMessagesResponse, *view.ErrorResponse) {
	peerDirectMessages, err := svc.getPeerDirectMessagesRepository.GetPeerDirectMessages(
		request.UserID, request.PeerID, request.Prev, request.Limit)
	if err != nil {
		return view.GetDirectMessagesResponse{}, &view.ErrorResponse{
			Code:     enum.InternalServerError,
			Location: "repository",
			Reason:   fmt.Sprintf("failed to get peer direct messages - %s", err.Error()),
		}
	}

	var cursorPosition int
	if len(peerDirectMessages) > 0 {
		last := peerDirectMessages[len(peerDirectMessages)-1]
		cursorPosition = last.ID
	}
	return view.GetDirectMessagesResponse{
		Messages:       peerDirectMessages,
		CursorPosition: cursorPosition,
	}, nil
}