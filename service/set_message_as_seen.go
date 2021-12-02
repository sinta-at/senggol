package service

import (
	"senggol/view"
)

type SetMessageAsSeenService struct{}

func (svc SetMessageAsSeenService) PostDirectMessages(request view.SetMessageAsSeenRequest) *view.ErrorResponse {
	return nil
}