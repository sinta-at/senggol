package pkg

import (
	"encoding/json"
	"net/http"
)

func ReadJSONRequest(request *http.Request, requestBody interface{}) error {
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(&requestBody)
	if err != nil {
		return err
	}
	return nil
}

func WriteJSONResponse(responseWriter http.ResponseWriter, statusCode int, contentType string, response interface{}) {
	if len(contentType) > 0 {
		responseWriter.Header().Add("Content-Type", "application/json")
	}
	responseWriter.WriteHeader(statusCode)
	_ = json.NewEncoder(responseWriter).Encode(response)
}