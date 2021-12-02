package view

import (
	"senggol/model"
)

type GetPeersResponse struct {
	Peers     []model.Peer `json:"peers"`
	PageNum   int          `json:"page_num"`
	PageSize  int          `json:"page_size"`
	PageCount int          `json:"page_count"`
}