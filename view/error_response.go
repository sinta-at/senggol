package view

import (
	"senggol/enum"
)

type ErrorResponse struct {
	Code     enum.ErrorCode `json:"code"`
	Location string         `json:"location"`
	Reason   string         `json:"reason"`
}