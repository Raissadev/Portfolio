package services

import (
	. "api/app/Repositories"
)

type UserService struct {
	ServiceInterface
	Repository UserRepository
}

func (us *UserService) All() string {
	return us.Repository.All()
}
