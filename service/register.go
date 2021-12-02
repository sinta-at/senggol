package service

import (
	"senggol/view"
)

type RegisterService struct{}

func (svc RegisterService) Register(request view.RegisterRequest) *view.ErrorResponse {
	return nil
}