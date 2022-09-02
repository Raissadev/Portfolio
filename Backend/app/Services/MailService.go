package services

import (
	. "api/app/Models"
	. "api/app/Repositories"
	"encoding/json"
)

type MailService struct {
	Repository MailRepository
	MailServiceInterface
}

type MailServiceInterface interface {
	Send(params *json.Decoder) (Mail, error)
}

func (us *MailService) Send(params *json.Decoder) (Mail, error) {
	return us.Repository.Send(params)
}
