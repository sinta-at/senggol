package service

import (
	"senggol/view"
)

type PostDirectMessageService struct{}

func (svc PostDirectMessageService) PostDirectMessage(request view.PostDirectMessageRequest) *view.ErrorResponse {
	return nil
}