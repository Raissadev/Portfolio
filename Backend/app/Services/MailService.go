package services

import (
	. "api/app/Models"
	. "api/app/Repositories"
	"encoding/json"
)

type MailService struct {
	ServiceInterface
	Repository MailRepository
}

func (us *MailService) Send(params *json.Decoder) (Mail, error) {
	return us.Repository.Send(params)
}
