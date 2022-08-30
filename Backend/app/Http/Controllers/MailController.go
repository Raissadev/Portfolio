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
	decoder := json.NewDecoder(r.Body)
	_, err := mailService.Send(decoder)

	if err != nil {
		uc.Message.WriteMessage(w, err.Error(), http.StatusInternalServerError)
		return
	}

	uc.Message.WriteMessage(w, "Email successfully sent!", http.StatusOK)
}
