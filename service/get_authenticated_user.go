package service

import (
	"fmt"
	"senggol/enum"
	"senggol/pkg"
	"senggol/repository"
	"senggol/view"
)

type GetAuthenticatedUserService struct {
	getUserByUsernameRepository repository.GetUserByUsername
}

func NewGetAuthenticatedUserService(getUserByUsernameRepository repository.GetUserByUsername) GetAuthenticatedUserService {
	return GetAuthenticatedUserService{getUserByUsernameRepository}
}

func (svc GetAuthenticatedUserService) GetAuthenticatedUser(request view.GetAuthenticatedUserRequest) (view.GetAuthenticatedUserResponse, *view.ErrorResponse) {
	user, err := svc.getUserByUsernameRepository.GetUserByUsername(request.Username)
	if err != nil {
		return view.GetAuthenticatedUserResponse{}, &view.ErrorResponse{
			Code:     enum.FailedAuthentication,
			Location: "service",
			Reason:   fmt.Sprintf("username or password is incorrect"),
		}
	}

	err = pkg.ValidateCredentials(request.Password, user.Password)
	if err != nil {
		return view.GetAuthenticatedUserResponse{}, &view.ErrorResponse{
			Code:     enum.FailedAuthentication,
			Location: "service",
			Reason:   fmt.Sprintf("username or password is incorrect"),
		}
	}

	return view.GetAuthenticatedUserResponse{user.ID}, nil
}