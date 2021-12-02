package service

import (
	"senggol/view"
)

type GetAuthenticatedUserService struct {}

func (svc GetAuthenticatedUserService) GetAuthenticatedUser(request view.GetAuthenticatedUserRequest) (view.GetAuthenticatedUserResponse, *view.ErrorResponse) {
	return view.GetAuthenticatedUserResponse{}, nil
}