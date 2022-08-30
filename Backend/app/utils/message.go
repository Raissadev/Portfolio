package utils

import (
	"encoding/json"
	"net/http"
)

type MessageRequest struct {
	//
}

type MessageRequestInterface interface {
	WriteMessage()
}

func (ms *MessageRequest) WriteMessage(w http.ResponseWriter, message string, status int) (int, error) {
	w.WriteHeader(status)

	data, err := json.Marshal(message)

	if err != nil {
		return 0, err
	}

	return w.Write(data)
}
