package service

import (
	"fmt"
	"math"
	"senggol/repository"
	"senggol/view"
)

type GetPeersService struct{
	getPeers      repository.GetPeers
	getPeersCount repository.GetPeersCount
}

func NewGetPeersService(getPeers repository.GetPeers, getPeersCount repository.GetPeersCount) GetPeersService {
	return GetPeersService{getPeers, getPeersCount}
}

func (svc GetPeersService) GetPeers(request view.GetPeersRequest) (view.GetPeersResponse, *view.ErrorResponse) {
	limit := request.PageSize
	offset := request.PageSize * (request.PageNum - 1)
	peers, err := svc.getPeers.GetPeers(request.UserID, limit, offset)
	if err != nil {
		return view.GetPeersResponse{}, &view.ErrorResponse{
			Code:     "INTERNAL_SERVER_ERROR",
			Location: "repository",
			Reason:   fmt.Sprintf("failed to get peers - %s", err.Error()),
		}
	}

	total, err := svc.getPeersCount.GetPeersCount(request.UserID)
	if err != nil {
		return view.GetPeersResponse{}, &view.ErrorResponse{
			Code:     "INTERNAL_SERVER_ERROR",
			Location: "repository",
			Reason:   fmt.Sprintf("failed to get peers count - %s", err.Error()),
		}
	}

	pageCount := int(math.Round(math.Ceil(float64(total) / float64(limit))))

	return view.GetPeersResponse{
		Peers:     peers,
		PageNum:   request.PageNum,
		PageSize:  request.PageSize,
		PageCount: pageCount,
	}, nil
}