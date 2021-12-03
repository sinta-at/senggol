package service

import (
	"fmt"
	"senggol/enum"
	"senggol/repository"
	"senggol/model"
	"senggol/pkg"
	"senggol/view"
	"time"
)

type RegisterService struct{
	createUserRepository    repository.CreateUser
	sequenceNumberGenerator pkg.SequenceNumberGenerator
}

func NewRegisterService(createUserRepository repository.CreateUser,
	sequenceNumberGenerator pkg.SequenceNumberGenerator) RegisterService {
	return RegisterService{createUserRepository, sequenceNumberGenerator}
}

func (svc RegisterService) Register(request view.RegisterRequest) *view.ErrorResponse {
	hashedPassword, err := pkg.GenerateCredentialsHash(request.Password)
	if err != nil {
		return &view.ErrorResponse{
			Code:     enum.InternalServerError,
			Location: "service",
			Reason:   "failed to construct user information",
		}
	}

	user := model.User{
		ID:        svc.sequenceNumberGenerator.Generate(),
		Username:  request.Username,
		Password:  hashedPassword,
		CreatedAt: time.Now(),
	}
	err = svc.createUserRepository.CreateUser(user)
	if err != nil {
		return &view.ErrorResponse{
			Code:     enum.InternalServerError,
			Location: "repository",
			Reason:   fmt.Sprintf("failed to create user - %s", err.Error()),
		}
	}

	return nil
}