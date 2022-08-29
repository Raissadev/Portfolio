package services

import (
	models "api/app/Models"
	. "api/app/Repositories"
)

type UserService struct {
	ServiceInterface
	Repository UserRepository
}

func (us *UserService) All() []models.User {
	return us.Repository.All()
}
