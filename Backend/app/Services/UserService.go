package services

import (
	. "api/app/Models"
	. "api/app/Repositories"
	"encoding/json"
)

type UserService struct {
	Repository UserRepository
	UserServiceInterface
}

type UserServiceInterface interface {
	All() []User
	Get(id uint64) (User, error)
	Create(params *json.Decoder) (User, error)
	Update(id uint64, params *json.Decoder) (User, error)
	Delete(id uint64) (bool, error)
}

func (us *UserService) All() []User {
	return us.Repository.All()
}

func (us *UserService) Get(id uint64) (User, error) {
	return us.Repository.Get(id)
}

func (us *UserService) Create(params *json.Decoder) (User, error) {
	return us.Repository.Create(params)
}

func (us *UserService) Update(id uint64, params *json.Decoder) (User, error) {
	return us.Repository.Update(id, params)
}

func (us *UserService) Delete(id uint64) (bool, error) {
	return us.Repository.Delete(id)
}
