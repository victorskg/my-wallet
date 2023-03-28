package response

import "encoding/json"

type messageResponse struct {
	Message string `json:"message"`
}

func MessageResponse(message string) []byte {
	messageBytes, _ := json.Marshal(&messageResponse{Message: message})
	return messageBytes
}
