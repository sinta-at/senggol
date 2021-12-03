package handler

import (
	"net/http"
	"senggol/enum"
	"senggol/pkg"
	"senggol/service"
	"senggol/view"

	"github.com/julienschmidt/httprouter"
)

type RegisterHandler struct {
	RegisterService service.RegisterService
}

func (handler RegisterHandler) HandleRegister(responseWriter http.ResponseWriter, request *http.Request, params httprouter.Params) {
	var registerRequest view.RegisterRequest
	err := pkg.ReadJSONRequest(request, &registerRequest)
	if err != nil {
		errorResponse := view.ErrorResponse{
			Code:     enum.BadRequest,
			Location: "controller",
			Reason:   "could not parse request",
		}
		pkg.WriteJSONResponse(responseWriter, enum.BadRequest.HttpStatusCode(), "application/json", errorResponse)
		return
	}

	errorResponse := handler.RegisterService.Register(registerRequest)
	if errorResponse != nil {
		pkg.WriteJSONResponse(responseWriter, errorResponse.Code.HttpStatusCode(), "application/json", errorResponse)
		return
	}

	responseWriter.WriteHeader(http.StatusCreated)
}