package service

import (
	"fmt"
	"senggol/enum"
	"senggol/model"
	"senggol/repository"
	"senggol/pkg"
	"senggol/view"
	"time"
)

type PostDirectMessageService struct {
	createDirectMessageRepository        repository.CreateDirectMessage
	messageSequenceNumberGenerator       pkg.SequenceNumberGenerator
	directMessageSequenceNumberGenerator pkg.SequenceNumberGenerator
}

func NewPostDirectMessageService(createDirectMessageRepository repository.CreateDirectMessage,
    messageSequenceNumberGenerator, directMessageSequenceNumberGenerator pkg.SequenceNumberGenerator) PostDirectMessageService {
    	return PostDirectMessageService{createDirectMessageRepository,
    	    messageSequenceNumberGenerator, directMessageSequenceNumberGenerator}
}

func (svc PostDirectMessageService) PostDirectMessage(request view.PostDirectMessageRequest) *view.ErrorResponse {
	now := time.Now()
	messageID := svc.messageSequenceNumberGenerator.Generate()

	message := model.Message{
		ID:        messageID,
		Content:   request.Content,
		CreatedAt: now,
	}
	inboundDirectMessage := model.DirectMessage{
		ID:        svc.directMessageSequenceNumberGenerator.Generate(),
		UserID:    request.PeerID,
		PeerID:    request.UserID,
		MessageID: messageID,
		Direction: "INBOUND",
	}
	outboundDirectMessage := model.DirectMessage{
		ID:        svc.directMessageSequenceNumberGenerator.Generate(),
		UserID:    request.UserID,
		PeerID:    request.PeerID,
		MessageID: messageID,
		Direction: "OUTBOUND",
	}

	err := svc.createDirectMessageRepository.CreateDirectMessage(message, inboundDirectMessage, outboundDirectMessage)
	if err != nil {
		return &view.ErrorResponse{
			Code:     enum.InternalServerError,
			Location: "repository",
			Reason:   fmt.Sprintf("failed to create direct message - %s", err.Error()),
		}
	}
	return nil
}