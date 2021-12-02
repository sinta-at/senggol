package service

import (
	"senggol/view"
)

type GetPeersService struct{}

func (svc GetPeersService) GetPeers(request view.GetPeersRequest) (view.GetPeersResponse, *view.ErrorResponse) {
	return view.GetPeersResponse{}, nil
}