package response

import (
	"encoding/json"
	"net/http"
)

func WriteResponseMessage(w http.ResponseWriter, message string, statusCode int) {
	w.WriteHeader(statusCode)
	w.Write(MessageResponse(message))
}

func WriteJSONResponse(w http.ResponseWriter, response any, statusCode int) {
	jsonBytes, _ := json.Marshal(response)
	w.WriteHeader(statusCode)
	w.Write(jsonBytes)
}
