package services

import (
	models "api/app/Models"
	. "api/app/Repositories"
	"encoding/json"
)

type UserService struct {
	ServiceInterface
	Repository UserRepository
}

func (us *UserService) All() []models.User {
	return us.Repository.All()
}

func (us *UserService) Get(id uint64) (models.User, error) {
	return us.Repository.Get(id)
}

func (us *UserService) Create(params *json.Decoder) (models.User, error) {
	return us.Repository.Create(params)
}
