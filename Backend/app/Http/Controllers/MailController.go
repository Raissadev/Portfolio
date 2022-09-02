package controllers

import (
	. "api/app/Services"
	"api/app/utils"
	"encoding/json"
	"net/http"
)

type MailController struct {
	Message *utils.MessageRequest
}

var mailService MailService

func (uc *MailController) Store(w http.ResponseWriter, r *http.Request) {
	_, err := mailService.Send(json.NewDecoder(r.Body))

	if err != nil {
		uc.Message.WriteMessage(w, "invalid data!", http.StatusNotAcceptable)
		return
	}

	uc.Message.WriteMessage(w, "Email successfully sent!", http.StatusOK)
}
