package service

import (
	"fmt"
	"senggol/enum"
	"senggol/repository"
	"senggol/view"
	"time"
)

type SetMessageSeenService struct {
	updateDirectMessageSeenAtRepository repository.UpdateDirectMessageSeenAt
}

func NewSetMessageSeenService(updateDirectMessageSeenAtRepository repository.UpdateDirectMessageSeenAt) SetMessageSeenService {
	return SetMessageSeenService{updateDirectMessageSeenAtRepository}
}

func (svc SetMessageSeenService) SetMessageSeen(request view.SetMessageSeenRequest) *view.ErrorResponse {
	err := svc.updateDirectMessageSeenAtRepository.UpdateDirectMessageSeenAt(request.MessageID, time.Now())
	if err != nil {
		return &view.ErrorResponse{
			Code:     enum.InternalServerError,
			Location: "repository",
			Reason:   fmt.Sprintf("failed to update direct message viewing status - %s", err.Error()),
		}
	}
	return nil
}