package service

import (
	"senggol/view"
)

type GetDirectMessagesService struct {}

func (svc GetDirectMessagesService) GetDirectMessages(request view.GetDirectMessagesRequest) (view.GetDirectMessagesResponse, *view.ErrorResponse) {
	return view.GetDirectMessagesResponse{}, nil
}