package handler

import (
	"fmt"
	"net/http"
	"senggol/enum"
	"senggol/pkg"
	"senggol/service"
	"senggol/view"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type GetDirectMessagesHandler struct {
	GetDirectMessagesService service.GetDirectMessagesService
}

func (handler GetDirectMessagesHandler) HandleGetDirectMessages(responseWriter http.ResponseWriter, request *http.Request, params httprouter.Params) {
	queryParams := request.URL.Query()

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

	prev := 0
	prevParam := queryParams["prev"]
	if len(prevParam) > 0 {
		prev, _ = strconv.Atoi(prevParam[0])
	}

	limitParam := queryParams["limit"]
	if len(limitParam) <= 0 {
		errorResponse := view.ErrorResponse{
			Code:     enum.BadRequest,
			Location: "controller",
			Reason:   "must specify limit",
		}
		pkg.WriteJSONResponse(responseWriter, enum.BadRequest.HttpStatusCode(), "application/json", errorResponse)
		return
	}

	limit, err := strconv.Atoi(limitParam[0])
	if !ok {
		errorResponse := view.ErrorResponse{
			Code:     enum.BadRequest,
			Location: "controller",
			Reason:   fmt.Sprintf("invalid limit - %s", err.Error()),
		}
		pkg.WriteJSONResponse(responseWriter, enum.BadRequest.HttpStatusCode(), "application/json", errorResponse)
		return
	}

	getDirectMessagesRequest := view.GetDirectMessagesRequest{
		UserID: userID,
		PeerID: peerID,
		Prev:   prev,
		Limit:  limit,
	}

	getDirectMessagesResponse, errorResponse := handler.GetDirectMessagesService.GetDirectMessages(getDirectMessagesRequest)
	if errorResponse != nil {
		pkg.WriteJSONResponse(responseWriter, errorResponse.Code.HttpStatusCode(), "application/json", errorResponse)
		return
	}

	pkg.WriteJSONResponse(responseWriter, http.StatusOK, "application/json", getDirectMessagesResponse)
}