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

type GetPeersHandler struct {
	GetPeersService service.GetPeersService
}

func (handler GetPeersHandler) HandleGetPeers(responseWriter http.ResponseWriter, request *http.Request, params httprouter.Params) {
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

	pageNumParam := queryParams["page_num"]
	if len(pageNumParam) <= 0 {
		errorResponse := view.ErrorResponse{
			Code:     enum.BadRequest,
			Location: "controller",
			Reason:   "must specify page_num",
		}
		pkg.WriteJSONResponse(responseWriter, enum.BadRequest.HttpStatusCode(), "application/json", errorResponse)
		return
	}

	pageNum, err := strconv.Atoi(pageNumParam[0])
	if err != nil {
		errorResponse := view.ErrorResponse{
			Code:     enum.BadRequest,
			Location: "controller",
			Reason:   fmt.Sprintf("invalid page_num - %s", err.Error()),
		}
		pkg.WriteJSONResponse(responseWriter, enum.BadRequest.HttpStatusCode(), "application/json", errorResponse)
		return
	}

	pageSizeParam := queryParams["page_size"]
	if len(pageSizeParam) <= 0 {
		errorResponse := view.ErrorResponse{
			Code:     enum.BadRequest,
			Location: "controller",
			Reason:   "must specify page_size",
		}
		pkg.WriteJSONResponse(responseWriter, enum.BadRequest.HttpStatusCode(), "application/json", errorResponse)
		return
	}

	pageSize, err := strconv.Atoi(pageSizeParam[0])
	if !ok {
		errorResponse := view.ErrorResponse{
			Code:     enum.BadRequest,
			Location: "controller",
			Reason:   fmt.Sprintf("invalid page_size - %s", err.Error()),
		}
		pkg.WriteJSONResponse(responseWriter, enum.BadRequest.HttpStatusCode(), "application/json", errorResponse)
		return
	}

	getPeersRequest := view.GetPeersRequest{
		UserID:   userID,
		PageNum:  pageNum,
		PageSize: pageSize,
	}

	getPeersResponse, errorResponse := handler.GetPeersService.GetPeers(getPeersRequest)
	if errorResponse != nil {
		pkg.WriteJSONResponse(responseWriter, errorResponse.Code.HttpStatusCode(), "application/json", errorResponse)
		return
	}

	pkg.WriteJSONResponse(responseWriter, http.StatusOK, "application/json", getPeersResponse)
}