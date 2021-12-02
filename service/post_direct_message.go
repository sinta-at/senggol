package service

import (
	"senggol/view"
)

type PostDirectMessagesService struct{}

func (svc PostDirectMessagesService) PostDirectMessages(request view.PostDirectMessagesRequest) *view.ErrorResponse {
	return nil
}