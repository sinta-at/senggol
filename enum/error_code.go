package enum

import (
	"net/http"
)

type ErrorCode string

const (
	BadRequest           ErrorCode = "BAD_REQUEST"
	InternalServerError  ErrorCode = "INTERNAL_SERVER_ERROR"
	FailedAuthentication ErrorCode = "FAILED_AUTHENTICATION"
)

func (ec ErrorCode) HttpStatusCode() int {
	switch ec {
	case BadRequest:
		return http.StatusBadRequest
	case InternalServerError:
		return http.StatusInternalServerError
	case FailedAuthentication:
		return http.StatusUnauthorized
	default:
		return http.StatusInternalServerError
	}
}