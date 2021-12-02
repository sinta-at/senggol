package view

type ErrorResponse struct {
	Code     string `json:"code"`
	Location string `json:"location"`
	Reason   string `json:"reason"`
}