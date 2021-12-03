package middleware

import (
	"context"
	"net/http"
	"senggol/enum"
	"senggol/pkg"
	"senggol/service"
	"senggol/view"

	"github.com/julienschmidt/httprouter"
)

type AuthMiddleware struct {
	GetAuthenticatedUserService service.GetAuthenticatedUserService
}

func (mw AuthMiddleware) BasicAuth(handler httprouter.Handle) httprouter.Handle {
	return func(responseWriter http.ResponseWriter, request *http.Request, params httprouter.Params) {
		username, password, isBasicAuth := request.BasicAuth()

		if !isBasicAuth {
			errorResponse := view.ErrorResponse{
				Code:     enum.FailedAuthentication,
				Location: "controller",
				Reason:   "unrecognized auth method",
			}
			pkg.WriteJSONResponse(responseWriter, enum.FailedAuthentication.HttpStatusCode(), "application/json", errorResponse)
			return
		}

		authRequest := view.GetAuthenticatedUserRequest{
			Username: username,
			Password: password,
		}
		authResponse, errorResponse := mw.GetAuthenticatedUserService.GetAuthenticatedUser(authRequest)
		if errorResponse != nil {
			pkg.WriteJSONResponse(responseWriter, errorResponse.Code.HttpStatusCode(), "application/json", errorResponse)
			return
		}

		updatedContext := context.WithValue(request.Context(), "UserID", authResponse.UserID)
		request = request.WithContext(updatedContext)
		handler(responseWriter, request, params)
	}
}