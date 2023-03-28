package response

import (
	"net/http"
)

func WriteResponseMessage(w http.ResponseWriter, message string, statusCode int) {
	w.WriteHeader(statusCode)
	w.Write(MessageResponse(message))
}
