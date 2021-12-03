package handler

import (
	"net/http"
	"senggol/enum"
	"senggol/pkg"
	"senggol/service"
	"senggol/view"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type PostDirectMessageHandler struct {
	PostDirectMessageService service.PostDirectMessageService
}

func (handler PostDirectMessageHandler) HandlePostDirectMessage(responseWriter http.ResponseWriter, request *http.Request, params httprouter.Params) {
	var postDirectMessageRequest view.PostDirectMessageRequest
	err := pkg.ReadJSONRequest(request, &postDirectMessageRequest)
	if err != nil {
		errorResponse := view.ErrorResponse{
			Code:     enum.BadRequest,
			Location: "controller",
			Reason:   "could not parse request",
		}
		pkg.WriteJSONResponse(responseWriter, enum.BadRequest.HttpStatusCode(), "application/json", errorResponse)
		return
	}

	contextUserID := request.Context().Value("UserID")
	if contextUserID == nil {
		errorResponse := view.ErrorResponse{
			Code:     enum.InternalServerError,
			Location: "controller",
			Reason:   "user ID not found in context",
		}
		pkg.WriteJSONResponse(responseWriter, enum.InternalServerError.HttpStatusCode(), "application/json", errorResponse)
		return
	}
	userID, ok := contextUserID.(int)
	if !ok {
		errorResponse := view.ErrorResponse{
			Code:     enum.InternalServerError,
			Location: "controller",
			Reason:   "invalid user ID in context",
		}
		pkg.WriteJSONResponse(responseWriter, enum.InternalServerError.HttpStatusCode(), "application/json", errorResponse)
		return
	}
	postDirectMessageRequest.UserID = userID

	peerIDParam := params.ByName("id")
	peerID, err := strconv.Atoi(peerIDParam)
	if err != nil {
		errorResponse := view.ErrorResponse{
			Code:     enum.BadRequest,
			Location: "controller",
			Reason:   "invalid peer id",
		}
		pkg.WriteJSONResponse(responseWriter, enum.BadRequest.HttpStatusCode(), "application/json", errorResponse)
		return
	}
	postDirectMessageRequest.PeerID = peerID

	errorResponse := handler.PostDirectMessageService.PostDirectMessage(postDirectMessageRequest)
	if errorResponse != nil {
		pkg.WriteJSONResponse(responseWriter, errorResponse.Code.HttpStatusCode(), "application/json", errorResponse)
		return
	}

	responseWriter.WriteHeader(http.StatusCreated)
}