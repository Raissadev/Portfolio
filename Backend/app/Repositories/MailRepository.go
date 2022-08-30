package repositories

import (
	. "api/app/Models"
	"encoding/json"
)

type MailRepository struct {
	Model Mail
	RepositoryInterface
}

func (r *MailRepository) Send(params *json.Decoder) (Mail, error) {
	return r.Model.Send(params)
}
