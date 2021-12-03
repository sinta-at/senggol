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

type SetMessageSeenHandler struct {
	SetMessageSeenService service.SetMessageSeenService
}

func (handler SetMessageSeenHandler) HandleSetMessageSeen(responseWriter http.ResponseWriter, request *http.Request, params httprouter.Params) {
	messageIDParam := params.ByName("id")
	messageID, err := strconv.Atoi(messageIDParam)
	if err != nil {
		errorResponse := view.ErrorResponse{
			Code:     enum.BadRequest,
			Location: "controller",
			Reason:   "invalid message id",
		}
		pkg.WriteJSONResponse(responseWriter, enum.BadRequest.HttpStatusCode(), "application/json", errorResponse)
		return
	}
	setMessageSeenRequest := view.SetMessageSeenRequest{
		MessageID: messageID,
	}

	errorResponse := handler.SetMessageSeenService.SetMessageSeen(setMessageSeenRequest)
	if errorResponse != nil {
		pkg.WriteJSONResponse(responseWriter, errorResponse.Code.HttpStatusCode(), "application/json", errorResponse)
		return
	}

	responseWriter.WriteHeader(http.StatusNoContent)
}