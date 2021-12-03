package service

import (
	"senggol/view"
)

type SetMessageSeenService struct{}

func (svc SetMessageSeenService) SetMessageSeen(request view.SetMessageSeenRequest) *view.ErrorResponse {
	return nil
}